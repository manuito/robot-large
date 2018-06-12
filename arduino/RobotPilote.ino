#include <SR04.h>
#include <LedControl.h>
#include <Adafruit_NeoPixel.h>
#include <Wire.h>
#include <pitches.h>

// Left HS sensor
#define HC_LEFT_ECHO_PIN 5
#define HC_LEFT_TRIG_PIN 4

// Right HS sensor
#define HC_RIGHT_ECHO_PIN 3
#define HC_RIGHT_TRIG_PIN 2

// Beeper
#define BUZZER_PIN 7

// LED Matrix
#define LED_MRX_DATA_IN 12
#define LED_MRX_CS 11
#define LED_MRX_CLK 10

// Motor driver 
#define LEFT_DIR_PIN   32
#define LEFT_DRV_PIN   8 // Need PWM
#define RIGHT_DIR_PIN  33
#define RIGHT_DRV_PIN  9 // Need PWM

// RGB band
#define RGB_PIN 6

// Status LED
#define LED_STARTED 31

// Various delay
int monitoringDelay = 200;
int actionDelay = 500;
int shortestDelay = 50;

// And loop control
int monitoringStep = 0;
int actionStep = 0;

void setup() {  

  // Start sensors / actions
  startAccel();
  startFace();
  startSerial();
  startBand();
  setupMotors();
  
  // Do a beep
  doBeep();
  
  // Startup led
  pinMode(LED_STARTED, OUTPUT);
  digitalWrite(LED_STARTED, HIGH);   // turn the LED on (HIGH is the voltage level)

  // Ready (TODO : raspb ACK)
  sendReady();
}

void loop() {

  if(checkCommand()){
    processCommand();
  }
  
  if(monitoringStep >= monitoringDelay){
    doMonitoring();
    monitoringStep = 0;
  } else {
    monitoringStep += shortestDelay;
  }
  
  if(actionStep >= actionDelay){
    doBandAction();
    doFaceAction();
    actionStep = 0;
  } else {
    actionStep += shortestDelay;
  }

  updateAccel();
  updateDistances();
  doBandShortAction();

  delay(shortestDelay);
}
