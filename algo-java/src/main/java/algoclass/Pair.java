package algoclass;

public class Pair<L, R> {
  public final L l;
  public final R r;

  public Pair(L l, R r) {
    this.l = l;
    this.r = r;
  }

  public static <L, R> Pair<L, R> from(L l, R r) {
    return new Pair<>(l, r);
  }
}
