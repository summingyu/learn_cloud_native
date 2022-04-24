package main

import (
	"context"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	klog "k8s.io/klog/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"httpserver/metrics"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.Handle("/metrics", promhttp.Handler())
	return mux
}

func main() {
	flag.Parse()
	defer klog.Flush()
	flag.Set("v", "4")
	flag.Set("logtostderr", "true")
	flag.Set("alsologtostderr", "true")
	klog.Info("Starting http server...")
	metrics.Register()
	server := &http.Server{
		Addr:    ":80",
		Handler: handler(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			klog.Fatal("ListenAndServe()=%+s", err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	osCall := <-ch
	klog.Infof("system call: %+v", osCall)
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxShutDown); err != nil {
		klog.Fatalf("server.Shutdown(ctxShutdown)=%+s", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 记录接收的request中的hander,并写入到response header中
	resp := make(map[string]string)
	for k, v := range r.Header {
		stringV := strings.Join(v, " ")
		w.Header().Set(k, stringV)
		// 使用io.WriteString之后,将不再设置header值
		// io.WriteString(w, fmt.Sprintf("%s=%s\n", k, stringV))
		resp[k] = stringV
	}

	// 2. 获取当前环境VERSION变量,写入respones header中
	VERSION, ex := os.LookupEnv("VERSION")
	if !ex {
		VERSION = "0"
		klog.Warning("VERSION not found, set ", VERSION)
	}
	w.Header().Add("Version", VERSION)
	resp["Version"] = VERSION

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	// 3. 获取客户端ip
	clientIP := getCurrentIP(r)
	klog.Infof("client IP: %s, status: %d\n", clientIP, http.StatusOK)
	// write只响应一次
	w.Write(jsonResp)
}

func getCurrentIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
	klog.Infof("healthz status: %d\n", http.StatusOK)
}
