import axios from 'axios'

class Api {
  config = { baseURL: process.env.ENDPOINT };

  async get (uri) {
    const response = await axios.get(`${this.config.baseURL}${uri}`)

    return response.data
  }

  async put (uri, body) {
    const response = await axios.put(`${this.config.baseURL}${uri}`, body)

    return response.data
  }
}

export default Api
