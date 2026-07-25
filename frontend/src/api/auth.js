const AUTH_URL = "http://localhost:8080";

function postForm(url, data) {
  return fetch(`${AUTH_URL}${url}`, {
    method: "POST",
    credentials: "include",
    body: new URLSearchParams(data),
  });
}

export async function login(email, password) {
  const res = await postForm("/login", {
    email,
    password,
  });

  if (!res.ok) {
    throw new Error(await res.text());
  }

  return res.json();
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
