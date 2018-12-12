#include <LedControl.h>
#include <Adafruit_NeoPixel.h>
#include <Wire.h>
#include <pitches.h>
#include <MPU6050_tockn_fork.h>

// Left HS sensor
#define HC_LEFT_ECHO_PIN    5
#define HC_LEFT_TRIG_PIN    4

// Right HS sensor
#define HC_RIGHT_ECHO_PIN   3
#define HC_RIGHT_TRIG_PIN   2

// Beeper
#define BUZZER_PIN          7

// LED Matrix
#define LED_MRX_DATA_IN     12
#define LED_MRX_CS          11
#define LED_MRX_CLK         10

// Motor driver 
#define LEFT_DIR_PIN        32
#define LEFT_DRV_PIN        8 // Need PWM
#define RIGHT_DIR_PIN       33
#define RIGHT_DRV_PIN       9 // Need PWM

// RGB band
#define RGB_PIN             6
#define RGB_LED_COUNT       15

// Status LED
#define LED_STARTED         31

// Various delay
#define MINIMAL_DELAY       25
#define MONITORING_DELAY    20 * MINIMAL_DELAY
#define SHORT_ACTION_DELAY  10 * MINIMAL_DELAY
#define ACTION_DELEAY       20 * MINIMAL_DELAY

// And loop control
int monitoringStep = 0;
int actionStep = 0;
int shortActionStep = 0;

void setup() {  

  pinMode(LED_STARTED, OUTPUT);
  digitalWrite(LED_STARTED, HIGH);
  
  // Start sensors / actions
  startFace();
  displayLoadingFace(1);
  delay(500);
  startSerial();
  pinMode(BUZZER_PIN, OUTPUT);
  startAccel();
  startBand();
  setupMotors();
  startHsSensors();
  displayLoadingFace(2);
  delay(500);
  zeroAccel();
  displayLoadingFace(3);
  
  // Do a beep
  doBeep();
  setFaceAction('1');
  
  // Ready (TODO : raspb ACK)
  sendReady();
  digitalWrite(LED_STARTED, LOW);
}

void loop() {

  if(checkCommand()){
    processCommand();
  }
  
  if(monitoringStep >= MONITORING_DELAY){
    updateDistances();
    updateAccel();
    doMonitoring();
    monitoringStep = 0;
  } else {
    monitoringStep += MINIMAL_DELAY;
  }
  
  if(actionStep >= ACTION_DELEAY){
    doBandAction();
    doFaceAction();
    doBuzzerAction();
    actionStep = 0;
  } else {
    actionStep += MINIMAL_DELAY;
  }

  if(actionStep >= SHORT_ACTION_DELAY){
    doBandShortAction();
    doBuzzerShortAction();
    shortActionStep = 0;
  } else {
    shortActionStep += MINIMAL_DELAY;
  }
  
  delay(MINIMAL_DELAY);
  digitalWrite(LED_STARTED, LOW);
}

