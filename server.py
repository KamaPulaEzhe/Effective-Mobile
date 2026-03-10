# """
# Простой веб-сервер на Python с несколькими страницами.
# Использует только стандартную библиотеку — никаких зависимостей.
# """

# from http.server import HTTPServer, BaseHTTPRequestHandler
# import json
# from datetime import datetime

# HOST = "127.0.0.1"
# PORT = 8000

# # ─── HTML-шаблоны ───────────────────────────────────────────────

# STYLE = """
# <style>
#   * { margin: 0; padding: 0; box-sizing: border-box; }
#   body { font-family: 'Segoe UI', Arial, sans-serif; background: #0f172a; color: #e2e8f0; }
#   nav { background: #1e293b; padding: 1rem 2rem; display: flex; gap: 1.5rem; align-items: center; }
#   nav a { color: #94a3b8; text-decoration: none; font-weight: 500; transition: color .2s; }
#   nav a:hover, nav a.active { color: #38bdf8; }
#   nav .logo { color: #38bdf8; font-weight: 700; font-size: 1.2rem; margin-right: auto; }
#   .container { max-width: 800px; margin: 3rem auto; padding: 0 2rem; }
#   h1 { font-size: 2.5rem; margin-bottom: 1rem; color: #f1f5f9; }
#   h2 { font-size: 1.5rem; margin-bottom: 1rem; color: #38bdf8; }
#   p { line-height: 1.8; color: #94a3b8; margin-bottom: 1rem; }
#   .card { background: #1e293b; border-radius: 12px; padding: 2rem; margin-bottom: 1.5rem;
#           border: 1px solid #334155; }
#   .card:hover { border-color: #38bdf8; }
#   .badge { display: inline-block; background: #0ea5e9; color: #fff; padding: .25rem .75rem;
#            border-radius: 999px; font-size: .85rem; margin-right: .5rem; }
#   footer { text-align: center; padding: 2rem; color: #475569; font-size: .85rem; }
#   ul { list-style: none; padding: 0; }
#   li { padding: .75rem 0; border-bottom: 1px solid #334155; }
#   li:last-child { border: none; }
#   code { background: #334155; padding: .15rem .4rem; border-radius: 4px; font-size: .9rem; }
# </style>
# """


# def nav(active="/"):
#     links = [("/", "Главная"), ("/about", "О проекте"), ("/api/time", "API")]
#     items = ""
#     for href, label in links:
#         cls = ' class="active"' if href == active else ""
#         items += f'<a href="{href}"{cls}>{label}</a>'
#     return f'<nav><span class="logo">🐍 PyServer</span>{items}</nav>'


# def page(title, body, active="/"):
#     return f"""<!DOCTYPE html>
# <html lang="ru">
# <head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
# <title>{title}</title>{STYLE}</head>
# <body>
# {nav(active)}
# <div class="container">{body}</div>
# <footer>Сделано на чистом Python · {datetime.now().year}</footer>
# </body></html>"""


# # ─── Страницы ────────────────────────────────────────────────────

# def home_page():
#     return page("Главная", """
#         <h1>Добро пожаловать!</h1>
#         <p>Это простой веб-сервер, написанный на чистом Python
#            без внешних зависимостей.</p>

#         <div class="card">
#             <h2>📄 Страницы</h2>
#             <ul>
#                 <li><a href="/" style="color:#38bdf8">/</a> — эта страница</li>
#                 <li><a href="/about" style="color:#38bdf8">/about</a> — о проекте</li>
#                 <li><a href="/api/time" style="color:#38bdf8">/api/time</a> — JSON API с текущим временем</li>
#             </ul>
#         </div>

#         <div class="card">
#             <h2>🛠 Технологии</h2>
#             <p>
#                 <span class="badge">Python 3</span>
#                 <span class="badge">http.server</span>
#                 <span class="badge">HTML/CSS</span>
#             </p>
#         </div>
#     """, active="/")


# def about_page():
#     return page("О проекте", """
#         <h1>О проекте</h1>
#         <div class="card">
#             <h2>Зачем?</h2>
#             <p>Показать, что для простого сервера не нужен Flask или Django.
#                Стандартная библиотека Python умеет всё, что нужно для прототипа.</p>
#         </div>
#         <div class="card">
#             <h2>Как запустить?</h2>
#             <p>Создай виртуальное окружение и запусти:</p>
#             <p><code>python3 -m venv venv</code></p>
#             <p><code>source venv/bin/activate</code></p>
#             <p><code>python server.py</code></p>
#         </div>
#         <div class="card">
#             <h2>Как остановить?</h2>
#             <p>Нажми <code>Ctrl+C</code> в терминале.</p>
#         </div>
#     """, active="/about")


# def not_found_page():
#     return page("404", """
#         <h1>404 — Страница не найдена 😕</h1>
#         <p>Попробуй вернуться на <a href="/" style="color:#38bdf8">главную</a>.</p>
#     """)


# # ─── Обработчик запросов ─────────────────────────────────────────

# class Handler(BaseHTTPRequestHandler):

#     def do_GET(self):
#         path = self.path.split("?")[0]  # убираем query-параметры

#         if path == "/":
#             self.respond(200, home_page())

#         elif path == "/about":
#             self.respond(200, about_page())

#         elif path == "/api/time":
#             # JSON API — возвращает текущее время
#             now = datetime.now()
#             data = {
#                 "time": now.strftime("%H:%M:%S"),
#                 "date": now.strftime("%Y-%m-%d"),
#                 "timestamp": int(now.timestamp()),
#             }
#             self.respond_json(200, data)

#         else:
#             self.respond(404, not_found_page())

#     # ── Вспомогательные методы ──

#     def respond(self, code, html):
#         self.send_response(code)
#         self.send_header("Content-Type", "text/html; charset=utf-8")
#         self.end_headers()
#         self.wfile.write(html.encode())

#     def respond_json(self, code, data):
#         self.send_response(code)
#         self.send_header("Content-Type", "application/json; charset=utf-8")
#         self.end_headers()
#         self.wfile.write(json.dumps(data, ensure_ascii=False, indent=2).encode())

#     def log_message(self, fmt, *args):
#         print(f"  {self.address_string()} → {args[0]}")


# # ─── Запуск ──────────────────────────────────────────────────────

# if __name__ == "__main__":
#     server = HTTPServer((HOST, PORT), Handler)
#     print(f"\n  🚀 Сервер запущен: http://{HOST}:{PORT}")
#     print(f"  Нажми Ctrl+C для остановки\n")
#     try:
#         server.serve_forever()
#     except KeyboardInterrupt:
#         print("\n  👋 Сервер остановлен")
#         server.server_close()






"""
Простой веб-сервер на Python — без HTML, только текст и JSON.
Стандартная библиотека, ноль зависимостей.
"""

from http.server import HTTPServer, BaseHTTPRequestHandler
import json
import random
from datetime import datetime
from urllib.parse import unquote_plus

HOST = "127.0.0.1"
PORT = 8000


# ─── Данные ──────────────────────────────────────────────────────

USERS = [
    {"id": 1, "name": "Алиса", "role": "admin"},
    {"id": 2, "name": "Борис", "role": "user"},
    {"id": 3, "name": "Вика", "role": "moderator"},
]

QUOTES = [
    "Простота — высшая форма изысканности. — Леонардо да Винчи",
    "Лучший код — ненаписанный код. — Джефф Этвуд",
    "Talk is cheap. Show me the code. — Линус Торвальдс",
    "Сначала реши проблему, потом пиши код. — Джон Джонсон",
]


# ─── Обработчик ──────────────────────────────────────────────────

class Handler(BaseHTTPRequestHandler):

    def do_GET(self):
        path = self.path.split("?")[0]

        if path == "/":
            self.text(200,
                "=== PyServer ===\n\n"
                "Доступные маршруты:\n\n"
                "  GET /           — эта справка\n"
                "  GET /time       — текущее время (JSON)\n"
                "  GET /users      — список пользователей (JSON)\n"
                "  GET /users/1    — пользователь по ID (JSON)\n"
                "  GET /quote      — случайная цитата (текст)\n"
                "  GET /echo?msg=  — эхо (JSON)\n"
            )

        elif path == "/time":
            now = datetime.now()
            self.json(200, {
                "date": now.strftime("%Y-%m-%d"),
                "time": now.strftime("%H:%M:%S"),
                "timestamp": int(now.timestamp()),
            })

        elif path == "/users":
            self.json(200, {"count": len(USERS), "users": USERS})

        elif path.startswith("/users/"):
            try:
                uid = int(path.split("/")[2])
                user = next((u for u in USERS if u["id"] == uid), None)
                if user:
                    self.json(200, user)
                else:
                    self.json(404, {"error": f"Пользователь {uid} не найден"})
            except ValueError:
                self.json(400, {"error": "ID должен быть числом"})

        elif path == "/quote":
            self.text(200, random.choice(QUOTES))

        elif path == "/echo":
            query = self.path.split("?", 1)[1] if "?" in self.path else ""
            params = dict(p.split("=", 1) for p in query.split("&") if "=" in p)
            params = {k: unquote_plus(v) for k, v in params.items()}
            self.json(200, {"echo": params})

        else:
            self.text(404, f"404 — маршрут «{path}» не найден")

    # ── Ответы ──

    def text(self, code, body):
        self.send_response(code)
        self.send_header("Content-Type", "text/plain; charset=utf-8")
        self.end_headers()
        self.wfile.write(body.encode())

    def json(self, code, data):
        self.send_response(code)
        self.send_header("Content-Type", "application/json; charset=utf-8")
        self.end_headers()
        self.wfile.write(json.dumps(data, ensure_ascii=False, indent=2).encode())

    def log_message(self, fmt, *args):
        print(f"  {self.address_string()} → {args[0]}")


# ─── Запуск ──────────────────────────────────────────────────────

if __name__ == "__main__":
    server = HTTPServer((HOST, PORT), Handler)
    print(f"\n  🚀 Сервер: http://{HOST}:{PORT}")
    print(f"  Ctrl+C — остановить\n")
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        print("\n  👋 Остановлен")
        server.server_close()