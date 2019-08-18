package main

import (
	pb "main/contract"
	"context"
	"google.golang.org/grpc"
	"log"
	"fmt"
	"time"
	"net/http"
)

func main() {

	http.HandleFunc("/",postMessage)
	err := http.ListenAndServe(":50050", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "<html>")
        fmt.Fprintln(w, "<head>")
        fmt.Fprintln(w, "<title>")
		fmt.Fprintln(w,"test")
		fmt.Fprintln(w, "</title>")
		fmt.Fprintln(w, "</head>")
		fmt.Fprintln(w, "<body>")
		fmt.Fprintln(w, "<form action='/' method='post'>")
		fmt.Fprintln(w, "契約内容:<input type='text' name='contract'><br><br>")
		fmt.Fprintln(w, "氏:<input type='text' name='lastname'><br><br>")
		fmt.Fprintln(w, "名:<input type='text' name='firstname'><br><br>")
		fmt.Fprintln(w, "電話番号:<input type='text' name='telnumber'><br><br>")
		fmt.Fprintln(w, "<input type='submit' value='ポスト'>")
		fmt.Fprintln(w, "</form>")
	    fmt.Fprintln(w, "</body>")
		fmt.Fprintln(w, "</html>")
	} else {
		//POST時にgRPCサーバーにリクエストを送信する。
		target := "server:50051"
		conn, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			log.Fatal("did not connect: %s", err)
		}
		defer conn.Close()
		client := pb.NewContractServiceClient(conn)
		contract := r.FormValue("msg")
		firstname := r.FormValue("msg")
		lastname := r.FormValue("msg")
		telnumber := r.FormValue("msg")
		ctx, cancel := context.WithTimeout(
			context.Background(), time.Second)
		defer cancel()
		result, err := client.Contract(ctx, &pb.ContractRequest{Contract: contract, Lastname: lastname,
											Firstname: firstname, Telnumber: telnumber})

		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, "err")
		}
		fmt.Fprintln(w, result.GetResult())

	}
}
