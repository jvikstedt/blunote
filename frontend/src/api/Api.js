import axios from 'axios'

class Api {
  baseURL = process.env.ENDPOINT
  config = {
    headers: {'Content-Type': 'application/json'}
  };

  async get (uri) {
    const response = await axios.get(`${this.baseURL}${uri}`)

    return response.data
  }

  async post (uri, body) {
    const response = await axios.post(`${this.baseURL}${uri}`, body, this.config)

    return response.data
  }

  async put (uri, body) {
    const response = await axios.put(`${this.baseURL}${uri}`, body, this.config)

    return response.data
  }
}

export default Api
