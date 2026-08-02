const REPORT_URL = "/api"


export function report(ip_address){
    return fetch(`${REPORT_URL}/report/`, {
        method:"POST",
        headers: {"Content-Type":"application/json"},
        body: JSON.stringify({ ip_address }),
    });
}
