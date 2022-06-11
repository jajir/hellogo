package main

/*
sudo apt install libusb-1.0-0-dev
sudo get install -y libusb-dev libusb-1.0-0-dev

brew install pkg-config
brew install libusb
*/

import (
	"log"

	libusb "github.com/gotmc/libusb/v2"
)

func main() {
	ctx, err := libusb.NewContext()
	if err != nil {
		log.Fatal("Couldn't create USB context. Ending now.")
	}
	defer ctx.Close()
	devices, err := ctx.DeviceList()
	if err != nil {
		log.Fatalf("Couldn't get devices")
	}
	log.Printf("Found %v USB devices.\n", len(devices))
	for _, device := range devices {
		usbDeviceDescriptor, err := device.DeviceDescriptor()
		if err != nil {
			log.Printf("Error getting device descriptor: %s", err)
			continue
		}
		handle, err := device.Open()
		if err != nil {
			log.Printf("Error opening device: %s", err)
			continue
		}
		defer handle.Close()
		serialNumber, err := handle.StringDescriptorASCII(usbDeviceDescriptor.SerialNumberIndex)
		if err != nil {
			serialNumber = "N/A"
		}
		manufacturer, err := handle.StringDescriptorASCII(usbDeviceDescriptor.ManufacturerIndex)
		if err != nil {
			manufacturer = "N/A"
		}
		product, err := handle.StringDescriptorASCII(usbDeviceDescriptor.ProductIndex)
		if err != nil {
			product = "N/A"
		}
		log.Printf("Found %s (S/N: %s) manufactured by %s",
			product,
			serialNumber,
			manufacturer,
		)
	}

}
