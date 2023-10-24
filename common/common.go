package common

import "change-api/dto"

// Logo 图形生成网站：http://patorjk.com/software/taag/
var Logo = `
 ▄████▄   ██░ ██  ▄▄▄       ███▄    █   ▄████ ▓█████    
▒██▀ ▀█  ▓██░ ██▒▒████▄     ██ ▀█   █  ██▒ ▀█▒▓█   ▀    
▒▓█    ▄ ▒██▀▀██░▒██  ▀█▄  ▓██  ▀█ ██▒▒██░▄▄▄░▒███      
▒▓▓▄ ▄██▒░▓█ ░██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█  ██▓▒▓█  ▄    
▒ ▓███▀ ░░▓█▒░██▓ ▓█   ▓██▒▒██░   ▓██░░▒▓███▀▒░▒████▒   
░ ░▒ ▒  ░ ▒ ░░▒░▒ ▒▒   ▓▒█░░ ▒░   ▒ ▒  ░▒   ▒ ░░ ▒░ ░   
  ░  ▒    ▒ ░▒░ ░  ▒   ▒▒ ░░ ░░   ░ ▒░  ░   ░  ░ ░  ░   
░         ░  ░░ ░  ░   ▒      ░   ░ ░ ░ ░   ░    ░      
░ ░       ░  ░  ░      ░  ░         ░       ░    ░  ░   
░                                                       
`

// Runtime 系统运行参数
var Runtime = dto.Runtime{
    Listen: "0.0.0.0",
    Port:   "8000",
    Config: "config/application.yaml",
}

// Developer 开发者信息
var Developer = dto.Developer{
    Name:  "Dylan",
    Email: "1214966109@qq.com",
}

// 版本信息
var Version = dto.Version{
    SystemVersion: "v1.0",
    GoVersion:     "1.20",
}
