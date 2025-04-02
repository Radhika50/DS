import java.util.Scanner;

public class Divisors {
  public static void main(String[] args) {
      int n;
      Scanner sc=new Scanner(System.in);
      n=sc.nextInt();
      int cnt=0;
      for(int i=1;i*i<=n;i++){
        if(n%i==0){
          cnt++;
          if(n/i != i){
            cnt++;
          }
        }
      }
      System.out.println("Divisor Cnt: "+cnt);
  }
}
