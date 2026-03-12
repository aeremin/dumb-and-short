import axios from 'axios';

const BACKEND_URL = 'https://europe-west6-dumb-and-short.cloudfunctions.net'

export async function create(url: string): Promise<string> {
  const response = await axios.post(BACKEND_URL + '/create', {url});
  return response.data.id;
}

export async function resolve(id: string): Promise<string> {
  const response = await axios.post(BACKEND_URL + '/resolve', {id});
  return response.data.url;
}