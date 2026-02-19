export async function generateNotes(content:string){return fetch("/api/notes",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({content})});}

export async function exportNotes(id:string){return fetch(`/api/notes/${id}/export`);}