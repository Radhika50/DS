import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Primefactors {
  public static void main(String[] args) {
      int n;
      Scanner sc=new Scanner(System.in);
      n=sc.nextInt();
      sc.close();
      List<Integer> fac=new ArrayList<>();
      for(int i=2;i*i<=n;i++){
        if(n%i==0){
          fac.add(i);
          while(n%i==0) {
            n=n/i;
          }
        }
      }
        if(n!=1) fac.add(n);
        System.out.println("Prime factors: "+fac);
  }
}
