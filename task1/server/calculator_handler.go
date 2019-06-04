package server

import (
	"context"
	"awesomeProject/task1/ayushpb/proto"
	"io"
	"log"
	//"time"
)

type CalculatorHandler struct{}

func (ch * CalculatorHandler) Tables(ctx context.Context, request *ayushpb.TableRequest) (*ayushpb.TableResponse, error) {
	response := &ayushpb.TableResponse{}
	response.Result = request.GetNumber() * request.GetNumber()*request.GetNumber()
	return response, nil
}

// Server Streaming
func (ch * CalculatorHandler) Greet(ctx context.Context,request *ayushpb.GreetingsRequest) (*ayushpb.GreetingsResponse, error) {

	response := &ayushpb.GreetingsResponse{}
	response.B = "Hi "+request.A+" you are doing really well"
	return response, nil
}
// Client Streaming

// BiDirectional Streaming
func (ch * CalculatorHandler) GCD(stream ayushpb.CalculatorService_GCDServer) error {
	nume:=[]int32{1}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatal(err)
		} else {
			number := req.GetNumber()
			nume=append(nume,number)
			result:=nume[0]
				for i:=1;i<len(nume);i++{
					result= int32(gcd(int(nume[i]), int(result)))
				}
			_ = stream.Send(&ayushpb.DivisorResponse{
				Num: result,
			})
			}
			}

		}
func gcd(number int,q int) int{
	if (number == 0) {
		return q;
	}
	return gcd(q%number, number);
}