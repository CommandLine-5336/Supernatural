const API_URL = "http://localhost:8585"

export function erase(){
    return fetch(`${API_URL}/erase`, {
        method:"POST",
    });
}
