package config

import "github.com/manifoldco/promptui"

const Logo string = `
   _____       _____                  
  / ____|     / ____|                 
 | |  __  ___| (___  _ __  _   _ _ __ 
 | | |_ |/ _ \\___ \| '_ \| | | | '__|
 | |__| | (_) |___) | |_) | |_| | |   
  \_____|\___/_____/| .__/ \__,_|_|   
                    | |               
                    |_|               
`

var LogoColoured string = promptui.Styler(promptui.FGCyan, promptui.FGBold)(Logo)
