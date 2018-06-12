// HS Sensors
SR04 sr04Left = SR04(HC_LEFT_ECHO_PIN,HC_LEFT_TRIG_PIN);
SR04 sr04Right = SR04(HC_RIGHT_ECHO_PIN,HC_RIGHT_TRIG_PIN);
long leftDist, rightDist;

void updateDistances(){
  leftDist = sr04Left.Distance();
  rightDist = sr04Right.Distance();
}

/**
 * For monitoring
 */
 
long getLeftDist(){
  return leftDist;
}

long getRightDist(){
  return rightDist;
}

