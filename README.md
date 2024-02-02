<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>




<!-- PROJECT SHIELDS -->
<!--
-->




<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">Line-bot-GPT聊天機器人</h3>

  <p align="center">
    A boyfriend chat robot that conencts to MongoDB
    <br />
  </p>
</div>

## Demo
https://www.youtube.com/watch?v=wC_4MDHLPck



<!-- ABOUT THE PROJECT -->
## About The Project

<a href="https://imgur.com/gZDPSj0"><img src="https://i.imgur.com/gZDPSj0.png" title="source: imgur.com" /></a>

A simple Line chat robot, using Line Messaging API and go-linebok-sdk that receive and save text messsages inside MongoDB docker and send back messages using GPT3.5-turbo model API
The Linbot name is called Boyfriend.ai
* The GPT3.5 turbo model has server personality set to "lovely Taiwanese boyfriend" with fast response time
* API call that can manully send message
* API query call that can retrieve all messages sent by a particular userID

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

Golang 1.20
</br>
MongoDB docker (version 5.0) 
</br>
git
</br>
postman 
</br>
ngrok 
</br>
third-party integration (Line https://developers.line.biz/en/docs/)  
docker

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

Run to following powershell script to set up mongodb docker (Window 11 Pro)
* ph1
  ```sh
  powershell -ExecutionPolicy Bypass -File mongodb-setup.ps1.
  ```
* in root file add '/chat/app/setting.yml', and insert channel secret and channel access token in
  ```
  channel_secret: "YOUR_CHANNEL_SECRET"
  channel_access_token: "YOUR_CHANNEL_ACCESS_TOKEN"
  ```
* go to pkg/gpt/gpt.go change line 48 to your own GPT API key
  ```
  req.Header.Set("Authorization", "Bearer <YOUR_API_KEY>")
  ```
* in root directory dowload all packages 
  ```
  go mod download
### Run

1. Start the backend server (running on PORT:8080
   ```sh
   go run main.go
   ```
2. activate ngrok and paste the URL onto LINE Developer Console underneath Message API / Webhook URL
   ```sh
   ngrok http 8080
   ```
3. Use the QR Code given to interact with LINE bot

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

* Chatting naturally just like a boyfriend
<a href="https://imgur.com/xkl1lgQ"><img src="https://i.imgur.com/xkl1lgQ.png" title="source: imgur.com" /></a>

* Calling Send Message API /api/sendmessage/:id to send a message of "在幹嘛?"
<a href="https://imgur.com/FxqwFwE"><img src="https://i.imgur.com/FxqwFwE.png" title="source: imgur.com" /></a>

* GET all of the messages sent by a particular user
* <a href="https://imgur.com/zEnvi5E"><img src="https://i.imgur.com/zEnvi5E.png" title="source: imgur.com" /></a>
