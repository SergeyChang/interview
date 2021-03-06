
#include <locale.h>
#include <stdio.h>
#include <time.h>
#include <stddef.h>
 
int 
main (int argc,char* argv[]){
   time_t currtime;
   struct tm *timer;
   char buffer[80];
 
   time( &currtime );
   timer = localtime( &currtime );
 
   printf("Locale is: %s\n", setlocale(LC_MESSAGES,""));
   strftime(buffer,80,"%c", timer );
   printf("Date is: %s\n", buffer);
 
  
   printf("Locale is: %s\n", setlocale(LC_MESSAGES, "en_US"));
   strftime(buffer,80,"%c", timer );
   printf("Date is: %s\n", buffer);
 
   return(0);
}
