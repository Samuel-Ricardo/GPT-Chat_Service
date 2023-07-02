# GPT-Chat_Service

<p align="center"> 
  <a href="https://fullcycle.com.br/" target="_blank">
    <img width="auto" src="https://miro.medium.com/v2/resize:fit:650/1*HkZDCVL7uUpYxOqpR5e_wg.png"/>
  </a> 
</p>

<h4 align="center" >🚀 🟨 Full Cycle Imerssion - 2023 🟨 🚀</h4>

<h4 align="center">
  Application developed during a Programmer Event, the <a style="color: #8a4af3;" href="https://github.com/search?q=imers%C3%A3o%20full%20cycle&type=repositories" target="_blank">Full Cycle Immersion</a> promoted by <a style="color: #8a4af3;" href="https://fullcycle.com.br/" target="_blank">@FullCycleSchool</a>
</h4>

#

<p align="center">
  |&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#project">Overview</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#techs">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#app">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#run-project">Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;
  <a style="color: #8a4af3;" href="#author">Author</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</p>

#

<h1 align="center">
  
  <a href="https://github.com/Samuel-Ricardo">
    <img src="https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=GITHUB"/>
  </a>

  <a herf="https://www.instagram.com/samuel_ricardo.ex/">
    <img src='https://img.shields.io/static/v1?label=&message=Samuel.ex&color=black&style=for-the-badge&logo=instagram'/> 
  </a>

  <a herf='https://www.linkedin.com/in/samuel-ricardo/'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=LinkedIn'/> 
  </a>

</h1>

<br>

<p id="project"/>

<h2>  | :artificial_satellite: About:  </h2>

<p align="justify">
    This project is one application of a complete Full Cycle Chat project to talk in real time with Chat-GPT, with amazing technologies like NextJS for front-end, Back-End for Front-end with NextJS, Next Auth + Keyclock for Authentication and Users management, MySQL database, GO Lang for microsservice, Docker and the power of Chat-GPT with real time reactive data and gRPC communication.
</p>

<p align="center">
  <img width="128px" src = "https://anch.ai/wp-content/uploads/2023/04/blog-new-chatgpt.png" />
  <img width="128px" src = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSbsRJ8KR7ADqtND-a1upt8M0wAufo6NBpYw31g68Eytz-9uXF5u32ziWo8JjeQf7OmNC0&usqp=CAU"/>
  <img width="128px" src = "https://www.mundodocker.com.br/wp-content/uploads/2015/06/docker_facebook_share.png" />
  <img width="128px" src = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSzKsNo97qflGoKifTQP6ztDT9tM-k2SIVsz7KV5vPsB0opSP00nLG6UIyy3oFSZDLkztA&usqp=CAU" />
  <img width="128px" src = "https://res.cloudinary.com/practicaldev/image/fetch/s--TpDTGYw5--/c_imagga_scale,f_auto,fl_progressive,h_900,q_auto,w_1600/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/yuxiptjqj8pa4bvyffym.png" />
  <img width="128px" src = "https://devtools.com.br/blog/wp-content/uploads/2013/06/MySQL-Logo.wine_.png" />
  <img width="128px" src = "https://grpc.io/img/logos/grpc-icon-color.png" /> 
  <img width="128px" src = "https://www.materialize.pro/wp-content/uploads/2021/10/GOLANG.png" /> 
</p>

<br>
<br>

<p align="justify">
  This microsservice propurse it's basically connect with Chat-GPT and handle with this data. Creating and managing Chats, receiving a message and calculating a response, all communication doesn't use REST & JSON's but gRCP that is more fast, light and economic. Using MySQL as DataBase.
</p>

> <a href="https://github.com/Samuel-Ricardo" target="_blank"> <img src="./readme_files/app_preview.png"> </a>

  <br>

#

<br/>

- Front-End & Back-End : NextJS | [ [repositories](https://github.com/Samuel-Ricardo/NextGPT) ]
- microsservice : GO Lang | [ [repositories](https://github.com/Samuel-Ricardo/GPT-Chat_Service) ]

#

<br>

<h2 id="techs">
  :building_construction: | Technologies and Concepts Studied:
</h2>

> <a href="https://openai.com/chatgpt" target="_blank"> <img width="128px" src = "https://anch.ai/wp-content/uploads/2023/04/blog-new-chatgpt.png" /> </a>

> <a href="https://go.dev/" target="_blank"> <img width="248px" src = "https://camo.githubusercontent.com/2b507540e2681c1a25698f246b9dca69c30548ed66a7323075b0224cbb1bf058/68747470733a2f2f676f6c616e672e6f72672f646f632f676f706865722f6669766579656172732e6a7067" /> </a>

<br>
<br>

- Go Lang
- Chat-GPT - [OpenAI]
- Docker
- gRPC
- Protobuf
- Performance
- MySQL
- Tiktoken
- Viper
- uuid
- Gotenv
- Chi
- Streams
- Real Time
- Multi Thread & Concurrency
- Scalability

> Among Others...

<br>

#

<h2 id="app">
  💻 | Application:
</h2>

> <a href="https://github.com/Samuel-Ricardo/NextGPT" target="_blank"> <img src="./readme_files/app_preview_chat.png" width="1280px"/> </a>

<p>
 This API open a webserver stream connection that receive chat status and datas in real time from Front-End, connect with Chat-GPT using OpenAI Client and send a "pre-message" that gives to GPT a little context of this chat, storage the messages and tokens in order, calculate the specific used model token limit and when hit the token limit i keep in other storage the old tokens, it gives to us a best control of GPT context and avoid unforeseens and extra expenses, with this in mund, the API connect with MySQL Database to create or update the current chat and messages. GPT-Chat will read all valid tokens and analize the chat context to give a response incrementally, so the API will receive from GPT and send to Front-End piece by piece of GPT response in real time
</p>

#

<h2 id="author">
  :octocat: | Author:  
</h2>

> <a target="_blank" href="https://www.linkedin.com/in/samuel-ricardo/"> <img width="350px" src="https://github.com/Samuel-Ricardo/bolao-da-copa/blob/main/readme_files/IMG_20220904_220148_188.jpg?raw=true"/> <br> <p> <b> - Samuel Ricardo</b> </p></a>

<h1>
  <a herf='https://github.com/Samuel-Ricardo'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=GITHUB'> 
  </a>
  
  <a herf='https://www.instagram.com/samuel_ricardo.ex/'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel.ex&color=black&style=for-the-badge&logo=instagram'> 
  </a>
  
  <a herf='https://twitter.com/SamuelR84144340'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=twitter'> 
  </a>
  
   <a herf='https://www.linkedin.com/in/samuel-ricardo/'>
    <img src='https://img.shields.io/static/v1?label=&message=Samuel%20Ricardo&color=black&style=for-the-badge&logo=LinkedIn'> 
  </a>
</h1>
