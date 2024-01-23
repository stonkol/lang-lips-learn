#include <stdint.h>

// Define addresses of relevant registers for GPIO and delay
#define GPIO_BASE_ADDRESS   0x40020000  // Example GPIO base address
#define GPIO_PORTA_DATA     (*(volatile uint32_t*)(GPIO_BASE_ADDRESS + 0x04))
#define GPIO_PORTA_DIR      (*(volatile uint32_t*)(GPIO_BASE_ADDRESS + 0x08))

#define DELAY_COUNT         1000000     // Adjust this value for the desired delay

void delay(uint32_t count) {
    while (count--);
}

int main() {
    // Set the GPIO pin as an output
    GPIO_PORTA_DIR |= (1 << 0);  // Assuming pin 0 is connected to the LED

    while (1) {
        // Toggle the LED state
        GPIO_PORTA_DATA ^= (1 << 0);

        // Delay to control the blink speed
        delay(DELAY_COUNT);
    }

    return 0;
}
