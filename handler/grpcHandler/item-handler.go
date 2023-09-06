package grpchandler

import (
	"context"
	"fmt"
	"io"
	"time"

	itempb "github.com/aom31/GO-Inventory/pkg/proto/proto-server"
	"github.com/aom31/GO-Inventory/src/service"
)

type itemGrpcHandler struct {
	itemService service.IItemService
}

func (itemGrpcHandler) mustEmbedUnimplementedItemServiceServer() {}

func (server *itemGrpcHandler) FindItems(stream itempb.ItemService_FindItemsServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	//0. เตรียม response ของ proto ออกไป
	itemsResponse := &itempb.ItemsList{
		Data: make([]*itempb.Item, 0),
	}

	//1. รับค่าจาก request client มาทำอะไรสักอย่าง
	// โดยจะเป็นการรับแบบ stream
	for {
		itemRequest, err := stream.Recv()
		//2. check ว่า client ส่งมาหมดยัง
		if err == io.EOF {
			fmt.Println("item_id out of range")
			break
		}
		if err != nil {
			return err
		}

		//3. ทำbuiss ไปดึง item จาก repo ออกมา เพื่อเตรียม response ไปให้ client
		item, err := server.itemService.FindOneItem(ctx, itemRequest.Id)
		if err != nil {
			return err
		}

		//4. เอา data ที่ได้เสร็จแล้ว มาใส่ใน model response ของ proto เพื่อตอบclint
		itemsResponse.Data = append(itemsResponse.Data, &itempb.Item{
			Id:          item.ObjectId.Hex(),
			Title:       item.Title,
			Description: item.Description,
			Damage:      item.Damage,
		})

	}

	//5. เมื่อ client ส่ง request มาครบ เราก็จะตอบกลับ response จาก server ออกไป

	return stream.SendAndClose(itemsResponse)
}
