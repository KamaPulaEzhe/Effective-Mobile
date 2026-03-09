# # import math

# # for n in range(10000, 1, -1):
# #     if math.ceil((math.log2(n) + 95)/8 + 40) * n <= 20 * 1024:
# #         print(n)
# #         break
    


# """
# Snakegram — Instagram-клон со змейкой.
# Один файл Python, ноль зависимостей.
# Запуск: python snakegram.py → http://127.0.0.1:8000
# """

# from http.server import HTTPServer, BaseHTTPRequestHandler

# HOST = "127.0.0.1"
# PORT = 8000

# HTML = r"""<!DOCTYPE html>
# <html lang="ru">
# <head>
# <meta charset="utf-8">
# <meta name="viewport" content="width=device-width,initial-scale=1,user-scalable=no">
# <title>Snakegram</title>
# <style>
# *{margin:0;padding:0;box-sizing:border-box}
# body{background:#000;color:#f5f5f5;font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Helvetica,Arial,sans-serif;overflow-x:hidden}
# ::-webkit-scrollbar{width:0}

# .app{max-width:470px;margin:0 auto;min-height:100vh;position:relative;padding-bottom:56px}

# /* Header */
# .header{position:sticky;top:0;z-index:100;background:#000;border-bottom:1px solid #262626;padding:10px 16px;display:flex;align-items:center;justify-content:space-between}
# .logo{font-family:'Segoe Script','Dancing Script',cursive;font-size:24px;font-weight:700;background:linear-gradient(45deg,#f09433,#e6683c,#dc2743,#cc2366,#bc1888);-webkit-background-clip:text;-webkit-text-fill-color:transparent;background-clip:text}
# .header-icons{display:flex;gap:20px;align-items:center}

# /* Stories */
# .stories{display:flex;gap:14px;padding:12px 16px;overflow-x:auto;border-bottom:1px solid #262626;scrollbar-width:none}
# .stories::-webkit-scrollbar{display:none}
# .story{display:flex;flex-direction:column;align-items:center;flex-shrink:0;cursor:pointer}
# .story-ring{width:66px;height:66px;border-radius:50%;padding:2.5px}
# .story-avatar{width:100%;height:100%;border-radius:50%;background:#000;border:2.5px solid #000;display:flex;align-items:center;justify-content:center;font-size:26px}
# .story-name{font-size:11px;color:#a8a8a8;margin-top:5px;max-width:64px;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}
# .story-add{position:absolute;bottom:-2px;right:-2px;width:20px;height:20px;border-radius:50%;background:#0095f6;border:2px solid #000;display:flex;align-items:center;justify-content:center;font-size:13px;font-weight:700;color:#fff}

# /* Post */
# .post{border-bottom:1px solid #262626}
# .post-header{display:flex;align-items:center;padding:10px 14px;gap:10px}
# .post-ava{width:34px;height:34px;border-radius:50%;background:linear-gradient(135deg,#f09433,#e6683c,#dc2743,#cc2366,#bc1888);padding:2px;flex-shrink:0}
# .post-ava-inner{width:100%;height:100%;border-radius:50%;background:#000;display:flex;align-items:center;justify-content:center;font-size:16px}
# .post-user{font-weight:600;font-size:13px}
# .post-meta{color:#a8a8a8;font-size:13px;margin-left:6px}
# .post-dots{color:#a8a8a8;font-size:18px;letter-spacing:2px;cursor:pointer;margin-left:auto}
# .post-actions{padding:10px 14px}
# .post-actions-row{display:flex;align-items:center;justify-content:space-between;margin-bottom:8px}
# .post-actions-left{display:flex;gap:16px;align-items:center}
# .post-actions-left svg{cursor:pointer}
# .post-likes{font-weight:600;font-size:13px;margin-bottom:4px}
# .post-caption{font-size:13px;line-height:1.4}
# .post-caption b{font-weight:600}
# .post-caption .more{color:#a8a8a8}
# .post-comments{color:#a8a8a8;font-size:12px;margin-top:6px}
# .post-time{color:#a8a8a8;font-size:10px;margin-top:4px;text-transform:uppercase;letter-spacing:.5px}

# /* Game */
# #game-wrap{position:relative;width:100%;aspect-ratio:1/1;background:#000}
# #game-canvas{display:block;width:100%;height:100%;image-rendering:pixelated}
# .game-overlay{position:absolute;inset:0;display:flex;flex-direction:column;align-items:center;justify-content:center;background:rgba(0,0,0,.7);cursor:pointer}
# .game-overlay h2{font-size:28px;font-weight:700;margin-bottom:4px}
# .game-overlay .score{color:#ed4956;font-size:18px;font-weight:600;margin-bottom:8px}
# .game-overlay p{color:#a8a8a8;font-size:13px}

# /* Feed images */
# .feed-img{width:100%;aspect-ratio:1/1;display:flex;align-items:center;justify-content:center;font-size:64px;opacity:.8}

# /* Heart animation */
# .heart-pop{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%) scale(0);font-size:80px;pointer-events:none;opacity:0}
# .heart-pop.show{animation:heartPop .8s ease forwards}
# @keyframes heartPop{0%{transform:translate(-50%,-50%) scale(0);opacity:1}15%{transform:translate(-50%,-50%) scale(1.2);opacity:1}30%{transform:translate(-50%,-50%) scale(1);opacity:1}70%{opacity:1}100%{opacity:0;transform:translate(-50%,-50%) scale(1)}}

# /* Bottom Nav */
# .bottom-nav{position:fixed;bottom:0;left:50%;transform:translateX(-50%);width:100%;max-width:470px;background:#000;border-top:1px solid #262626;display:flex;justify-content:space-around;align-items:center;padding:8px 0 14px;z-index:100}
# .bottom-nav svg{cursor:pointer;opacity:.6;padding:4px 12px}
# .bottom-nav svg:first-child{opacity:1}
# .nav-profile{cursor:pointer;width:26px;height:26px;border-radius:50%;border:1px solid #a8a8a8;display:flex;align-items:center;justify-content:center;font-size:14px}
# </style>
# </head>
# <body>
# <div class="app">

#   <!-- Header -->
#   <div class="header">
#     <div class="logo">Snakegram</div>
#     <div class="header-icons">
#       <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/></svg>
#       <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
#     </div>
#   </div>

#   <!-- Stories -->
#   <div class="stories" id="stories"></div>

#   <!-- Game Post -->
#   <div class="post">
#     <div class="post-header">
#       <div class="post-ava"><div class="post-ava-inner">🐍</div></div>
#       <div><span class="post-user">snake_game</span><span class="post-meta">• Play now</span></div>
#       <span class="post-dots">•••</span>
#     </div>
#     <div id="game-wrap">
#       <canvas id="game-canvas"></canvas>
#       <div class="game-overlay" id="game-overlay" onclick="startGame()">
#         <h2>🐍 Snake</h2>
#         <p>Tap or press arrow keys to play</p>
#       </div>
#       <div class="heart-pop" id="heart-pop">❤️</div>
#     </div>
#     <div class="post-actions">
#       <div class="post-actions-row">
#         <div class="post-actions-left">
#           <svg id="like-btn" onclick="toggleLike('game')" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/></svg>
#           <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
#           <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M22 2L11 13M22 2l-7 20-4-9-9-4z"/></svg>
#         </div>
#         <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
#       </div>
#       <div class="post-likes" id="game-likes">4,821 likes</div>
#       <div class="post-caption"><b>snake_game</b> Use arrow keys or swipe to play! 🎮 Best: <span id="best-score">0</span><span class="more"> ...more</span></div>
#       <div class="post-comments">View all 283 comments</div>
#       <div class="post-time">1 hour ago</div>
#     </div>
#   </div>

#   <!-- Feed -->
#   <div id="feed"></div>

#   <!-- Bottom Nav -->
#   <div class="bottom-nav">
#     <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/></svg>
#     <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M11 19a8 8 0 1 0 0-16 8 8 0 0 0 0 16zM21 21l-4.35-4.35"/></svg>
#     <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M12 5v14M5 12h14"/></svg>
#     <svg width="24" height="24" viewBox="0 0 24 24" fill="#f5f5f5" stroke="none"><path d="M7 2v20l10-10z"/></svg>
#     <div class="nav-profile">😎</div>
#   </div>
# </div>

# <script>
# /* ── Data ── */
# const STORIES=[
#   {user:"snake_master",emoji:"🐍",grad:"linear-gradient(135deg,#f09433,#ee2a7b,#6228d7)"},
#   {user:"pixel_queen",emoji:"👑",grad:"linear-gradient(135deg,#00c6ff,#0072ff)"},
#   {user:"code_ninja",emoji:"🔥",grad:"linear-gradient(135deg,#f857a6,#ff5858)"},
#   {user:"game_dev_42",emoji:"🎮",grad:"linear-gradient(135deg,#11998e,#38ef7d)"},
#   {user:"retro_vibes",emoji:"✨",grad:"linear-gradient(135deg,#fc5c7d,#6a82fb)"},
#   {user:"dev.js",emoji:"💻",grad:"linear-gradient(135deg,#ffd200,#f7971e)"},
#   {user:"css_wizard",emoji:"🎨",grad:"linear-gradient(135deg,#a18cd1,#fbc2eb)"},
# ];
# const POSTS=[
#   {user:"snake_master",avatar:"🐍",likes:2847,caption:"New high score today! 🔥 #SnakeGame",time:"2h",grad:"linear-gradient(135deg,#1a1a2e,#16213e,#0f3460)"},
#   {user:"pixel_queen",avatar:"👑",likes:1293,caption:"Who else is addicted? 😂",time:"4h",grad:"linear-gradient(135deg,#0f0c29,#302b63,#24243e)"},
#   {user:"game_dev_42",avatar:"🎮",likes:892,caption:"Built this in one file 💪 #coding",time:"6h",grad:"linear-gradient(135deg,#141e30,#243b55)"},
# ];
# const likes={};

# /* ── Stories ── */
# document.getElementById("stories").innerHTML=`
#   <div class="story">
#     <div class="story-ring" style="background:#262626;position:relative">
#       <div class="story-avatar">🎮</div>
#       <div class="story-add">+</div>
#     </div>
#     <span class="story-name">Your story</span>
#   </div>
# `+STORIES.map(s=>`
#   <div class="story">
#     <div class="story-ring" style="background:${s.grad}">
#       <div class="story-avatar">${s.emoji}</div>
#     </div>
#     <span class="story-name">${s.user}</span>
#   </div>
# `).join("");

# /* ── Feed ── */
# document.getElementById("feed").innerHTML=POSTS.map((p,i)=>`
#   <div class="post">
#     <div class="post-header">
#       <div class="post-ava" style="background:${STORIES[i]?.grad||'#262626'}"><div class="post-ava-inner">${p.avatar}</div></div>
#       <div><span class="post-user">${p.user}</span></div>
#       <span class="post-dots">•••</span>
#     </div>
#     <div class="feed-img" style="background:${p.grad}">${p.avatar}</div>
#     <div class="post-actions">
#       <div class="post-actions-row">
#         <div class="post-actions-left">
#           <svg id="like-btn-${i}" onclick="toggleLike(${i})" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/></svg>
#           <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>
#           <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M22 2L11 13M22 2l-7 20-4-9-9-4z"/></svg>
#         </div>
#         <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#f5f5f5" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"><path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"/></svg>
#       </div>
#       <div class="post-likes" id="likes-${i}">${p.likes.toLocaleString()} likes</div>
#       <div class="post-caption"><b>${p.user}</b> ${p.caption}</div>
#       <div class="post-time">${p.time}</div>
#     </div>
#   </div>
# `).join("");

# /* ── Likes ── */
# function toggleLike(id){
#   likes[id]=!likes[id];
#   if(id==="game"){
#     const btn=document.getElementById("like-btn");
#     const n=4821+(likes[id]?1:0);
#     document.getElementById("game-likes").textContent=n.toLocaleString()+" likes";
#     btn.setAttribute("fill",likes[id]?"#ed4956":"none");
#     btn.setAttribute("stroke",likes[id]?"#ed4956":"#f5f5f5");
#   } else {
#     const btn=document.getElementById("like-btn-"+id);
#     const n=POSTS[id].likes+(likes[id]?1:0);
#     document.getElementById("likes-"+id).textContent=n.toLocaleString()+" likes";
#     btn.setAttribute("fill",likes[id]?"#ed4956":"none");
#     btn.setAttribute("stroke",likes[id]?"#ed4956":"#f5f5f5");
#   }
# }

# /* ── Snake Game ── */
# const CELL=16,COLS=22,ROWS=22;
# const canvas=document.getElementById("game-canvas");
# canvas.width=COLS*CELL; canvas.height=ROWS*CELL;
# const ctx=canvas.getContext("2d");
# const overlay=document.getElementById("game-overlay");
# const DIR={UP:[0,-1],DOWN:[0,1],LEFT:[-1,0],RIGHT:[1,0]};
# let snake,dir,nextDir,food,score,bestScore=0,gameLoop=null,state="idle";

# function initSnake(){
#   const cx=COLS/2|0, cy=ROWS/2|0;
#   snake=[[cx,cy],[cx-1,cy],[cx-2,cy]];
#   dir=DIR.RIGHT; nextDir=DIR.RIGHT;
#   food=spawnFood(); score=0;
# }
# function spawnFood(){
#   let p; do{p=[Math.random()*COLS|0,Math.random()*ROWS|0]}
#   while(snake.some(s=>s[0]===p[0]&&s[1]===p[1])); return p;
# }

# function draw(){
#   ctx.fillStyle="#000"; ctx.fillRect(0,0,canvas.width,canvas.height);
#   // grid
#   ctx.fillStyle="#1a1a1a";
#   for(let x=0;x<COLS;x++) for(let y=0;y<ROWS;y++)
#     ctx.fillRect(x*CELL+CELL/2-.5,y*CELL+CELL/2-.5,1,1);
#   // food glow
#   const grd=ctx.createRadialGradient(food[0]*CELL+CELL/2,food[1]*CELL+CELL/2,2,food[0]*CELL+CELL/2,food[1]*CELL+CELL/2,CELL*2);
#   grd.addColorStop(0,"rgba(237,73,86,0.4)"); grd.addColorStop(1,"rgba(237,73,86,0)");
#   ctx.fillStyle=grd; ctx.fillRect(food[0]*CELL-CELL,food[1]*CELL-CELL,CELL*3,CELL*3);
#   // food
#   ctx.fillStyle="#ed4956"; ctx.shadowColor="#ed4956"; ctx.shadowBlur=8;
#   ctx.beginPath(); ctx.arc(food[0]*CELL+CELL/2,food[1]*CELL+CELL/2,CELL/2-2,0,Math.PI*2); ctx.fill();
#   ctx.shadowBlur=0;
#   // snake
#   snake.forEach(([x,y],i)=>{
#     const t=i/snake.length;
#     ctx.fillStyle=`rgb(${131+(88-131)*t|0},${58+(81-58)*t|0},${180+(219-180)*t|0})`;
#     const pad=i===0?1:2;
#     ctx.beginPath(); ctx.roundRect(x*CELL+pad,y*CELL+pad,CELL-pad*2,CELL-pad*2,4); ctx.fill();
#   });
#   // eyes
#   if(snake.length){
#     const[hx,hy]=snake[0]; ctx.fillStyle="#fff";
#     const ex=hx*CELL+CELL/2+dir[0]*2,ey=hy*CELL+CELL/2+dir[1]*2;
#     ctx.beginPath(); ctx.arc(ex-2,ey-1,1.5,0,Math.PI*2); ctx.arc(ex+2,ey+1,1.5,0,Math.PI*2); ctx.fill();
#   }
#   if(state==="playing"){
#     ctx.fillStyle="#fff"; ctx.font="bold 14px sans-serif"; ctx.textAlign="right";
#     ctx.fillText(score,canvas.width-10,20); ctx.textAlign="start";
#   }
# }

# function tick(){
#   dir=nextDir;
#   const[hx,hy]=snake[0]; const nx=hx+dir[0],ny=hy+dir[1];
#   if(nx<0||nx>=COLS||ny<0||ny>=ROWS||snake.some(s=>s[0]===nx&&s[1]===ny)){
#     clearInterval(gameLoop); state="over";
#     if(score>bestScore){bestScore=score; document.getElementById("best-score").textContent=bestScore}
#     overlay.innerHTML=`<h2>Game Over</h2><div class="score">Score: ${score}</div><p>Tap or press arrow to retry</p>`;
#     overlay.style.display="flex"; draw(); return;
#   }
#   snake.unshift([nx,ny]);
#   if(nx===food[0]&&ny===food[1]){score++;food=spawnFood()}else{snake.pop()}
#   draw();
# }

# function startGame(){
#   initSnake(); state="playing"; overlay.style.display="none";
#   draw(); clearInterval(gameLoop); gameLoop=setInterval(tick,120);
# }

# /* Keyboard */
# document.addEventListener("keydown",e=>{
#   if(state!=="playing"){
#     if(["ArrowUp","ArrowDown","ArrowLeft","ArrowRight"," "].includes(e.key)){e.preventDefault();startGame();return}
#   }
#   if(state!=="playing")return;
#   switch(e.key){
#     case"ArrowUp":case"w":if(dir!==DIR.DOWN)nextDir=DIR.UP;break;
#     case"ArrowDown":case"s":if(dir!==DIR.UP)nextDir=DIR.DOWN;break;
#     case"ArrowLeft":case"a":if(dir!==DIR.RIGHT)nextDir=DIR.LEFT;break;
#     case"ArrowRight":case"d":if(dir!==DIR.LEFT)nextDir=DIR.RIGHT;break;
#   }
#   if(["ArrowUp","ArrowDown","ArrowLeft","ArrowRight"].includes(e.key))e.preventDefault();
# });

# /* Touch */
# let ts=null;
# canvas.addEventListener("touchstart",e=>{ts={x:e.touches[0].clientX,y:e.touches[0].clientY}});
# canvas.addEventListener("touchend",e=>{
#   if(!ts)return;
#   const dx=e.changedTouches[0].clientX-ts.x, dy=e.changedTouches[0].clientY-ts.y;
#   if(Math.abs(dx)<20&&Math.abs(dy)<20){if(state!=="playing")startGame();return}
#   if(state!=="playing")return;
#   if(Math.abs(dx)>Math.abs(dy)){
#     if(dx>0&&dir!==DIR.LEFT)nextDir=DIR.RIGHT;
#     else if(dx<0&&dir!==DIR.RIGHT)nextDir=DIR.LEFT;
#   }else{
#     if(dy>0&&dir!==DIR.UP)nextDir=DIR.DOWN;
#     else if(dy<0&&dir!==DIR.DOWN)nextDir=DIR.UP;
#   }
# });

# /* Double-tap like */
# let lastTap=0;
# document.getElementById("game-wrap").addEventListener("click",()=>{
#   const now=Date.now();
#   if(now-lastTap<300&&state==="playing"){
#     likes["game"]=true; toggleLike("game");
#     const hp=document.getElementById("heart-pop");
#     hp.classList.remove("show"); void hp.offsetWidth; hp.classList.add("show");
#   }
#   lastTap=now;
# });

# initSnake(); draw();
# </script>
# </body>
# </html>
# """


# class Handler(BaseHTTPRequestHandler):
#     def do_GET(self):
#         self.send_response(200)
#         self.send_header("Content-Type", "text/html; charset=utf-8")
#         self.end_headers()
#         self.wfile.write(HTML.encode())

#     def log_message(self, fmt, *args):
#         print(f"  {self.address_string()} -> {args[0]}")


# if __name__ == "__main__":
#     server = HTTPServer((HOST, PORT), Handler)
#     print(f"\n  🐍 Snakegram: http://{HOST}:{PORT}")
#     print(f"  Ctrl+C — остановить\n")
#     try:
#         server.serve_forever()
#     except KeyboardInterrupt:
#         print("\n  👋 Остановлен")
#         server.server_close()




from ipaddress import *

cnt = 0
ip_house = ip_network("252.67.33.87/255.248.0.0", 0)
for ip_sasha in ip_house:
    s = f"{ip_sasha:b}"
    if s[16:].count("1") / s[:16].count("1") > 2:
        cnt += 1
        
print(cnt)


