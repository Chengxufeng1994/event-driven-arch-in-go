const axios = require('axios');

class Client {
  constructor(host = 'http://localhost:8080') {
    this.host = host;
  }

  startBasket(customerId) {
    return axios.post(`${this.host}/api/v1/baskets`, { customerId });
  }

  addItem(basketId, productId, quantity = 1) {
    return axios.put(`${this.host}/api/v1/baskets/${basketId}/addItem`, {
      productId,
      quantity,
    });
  }
}

module.exports = { Client };
