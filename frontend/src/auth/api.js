const AUTH_URL = "http://localhost:8080";

function postForm(url, data) {
  return fetch(`${AUTH_URL}${url}`, {
    method: "POST",
    credentials: "include",
    body: new URLSearchParams(data),
  });
}

export function login(email, password) {
  return postForm("/login", {
    email,
    password,
  });
}

export function register(email, password) {
  return postForm("/register", {
    email,
    password,
  });
}

export function logout() {
  return fetch(`${AUTH_URL}/logout`, {
    credentials: "include",
  });
}

export function session() {
  return fetch(`${AUTH_URL}/session`, {
    credentials: "include",
  });
}
