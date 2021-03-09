# ASCII ART
ASCII art is a type of graphic design that uses the 95 printable characters from the ASCII standard.
One of the first use cases of ASCII art was in early emails, when images could not yet be embedded.

## Golang ASCII Art
We use an image of the Golang mascotte and transform into ASCII art. To do that we first transform the RGB values into grayscale values and then divide them into 13 buckets depending on the darkness of the pixel.
Each of these pixel will then be assigned a symbols: " ", ".", ",", "-", "~", ":", ";", "=", "!", "\*", "#", "$", "@". These symbols are sorted from less dark to darkest and assigned to each pixels.

## Pattern finding
In order to find a certain list of symbols in the ASCII art and print them in another color we create a new structure Pattern. This structure will take as parameter the list of symbols and also the orientation of the pattern (vertical/horizontal).
Once the pattern struct is created we can define a function findPatternLocationthat goes through the ASCII art and outputs a 2d slice of pattern symbols postiion.
The 2d array can then be used to print the symbols into a different color using printASCIImageColorizedPattern.

## Possible improvements
- Be able to give as input a list of vertical pattern and a list of horizontal pattern in order to colorize 2d arrays of symbols rather than only 1d lines.


                                 .,,-~~-,                                  
                           .,-:;!====!====;;-. -=!!!:-.                    
                         -;=!*#=;!*@@@@@$#!:#=::=:==**=.                   
                       ~!*****:;@@@@@@@@@@@$-**=.., =**:                   
                    .:*******=-@@@@@@@@@=..::;***~ .***:                   
                  ,;*********-@@@@@@@@@$   ~--****!~;**                    
                -=***********-@@@@@@@@@@-  ~!,******=;-                    
              ~!;:;==;;;*****~$@@@@@@@@@@$$$@-********;                    
            ,!;~#@@@@@@@=;****~@@@@@@@@@@@@@#=*********;                   
           :*:=@@@@@@@$***:***==$@@@@@@@@@@*:***********;                  
          =*~$@@@@@@@*,  ~;~***==!#@@@@@#*!=*************~                 
         :*!-@@@@@@@@=  ,;*:******;~:;;:=*****************                 
        .**!~@@@@@@@@@~..~@=*!~,.-=***********************;                
     -;!-***-@@@@@@@@@@@@@@:;   . -::~:********************,               
    ;**#~***;#@@@@@@@@@@@@!;,  .-;====;~*******************=               
   ,**!~-****~#@@@@@@@@@@=;;:=========-=********************~              
   =**  ~*****;;#@@@@@@*~;*-=====~~=:-***********************.             
   ;**=--*******=;======***~;;;~!@=@@!;**********************!.            
   .=**#~******************!===:$@$*@@=***********************!.           
     -;!:***********************=@@*:==************************=           
        .***********************;=!==***************************,          
         :*******************************************************-!##*:    
          !******************************************************:@@@@=    
          ,******************************************************=,~!!,    
           :******************************************************         
           .!*****************************************************.        
            ,************************************!!!**************~        
             -**********************************~*!!=;************!        
              ~*********************************=@@@@@:************.       
               :!*****************************!=!##$@;:************.       
                -!***************************=:=!!!=!==************,       
                 .!***************************!********************~       
                  -************************************************:       
                   ;***********************************************~       
                   .***********************************************~       
                    !**********************************************:       
                    .**********************************************:       
                     =*********************************************~       
                     :*********************************************,       
                     ,*********************************************.       
                     .********************************************!        
                     .********************************************=        
                     ,********************************************-        
                     .********************************************         
                     .*******************************************;         
                      *******************************************~;~,      
                     ,******************************************:!@@$!     
                     ,*****************************************;.$@@@:.    
                     .****************************************:    #@:     
                      :*************************************;,      ,      
                      ~!**********************************;-               
                     :@~!*******************************:.                 
                        .!**************************!=-                    
                          ;!********************!=:-.                      
                           ,~!==;;**********=;~,.                          
                              =@@$:;*===;~,.                               
                             ~@@@@=                                        
                            ,@@@@;                                         
                            *;@@!                                          
                            .-,-                                           