package util;

final public class Arr {

  public static void swap(int[] x, int a, int b) {
    int t = x[a];
    x[a] = x[b];
    x[b] = t;
  }

}
