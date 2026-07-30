const CLEANUP_URL = "http://localhost:8585"

export function erase(){
    return fetch(`${CLEANUP_URL}/erase`, {
        method:"POST",
    });
}
