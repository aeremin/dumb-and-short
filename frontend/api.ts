import axios from 'axios';

const BACKEND_URL = 'https://dumb-and-short-815779847222.europe-west6.run.app'

export async function create(url: string) {
  const response = await axios.post(BACKEND_URL + '/create', {url});
  return response.data;
}

export async function resolve(id: string): Promise<string> {
  const response = await axios.post(BACKEND_URL + '/resolve', {id});
  return response.data.url;
}