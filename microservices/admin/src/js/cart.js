const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, Cart, CartProduct} = require('./protobuf/cart_pb.js');

const {CartAPIClient} = require('./protobuf/cart_grpc_web_pb.js');

export const order = new Vue({
  el: '#cart',
  data: {
    endpoint: window.location.protocol + '//' + window.location.host + "/cart",
    form: {
      userUUID: '',
      cartProducts: [],
    },
    resp: {
      cart: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new CartAPIClient(this.endpoint);
  },
  methods: {
    addCartProduct: function() {
      this.form.cartProducts.push({value:''});
    },
    clearForm: function() {
      this.form.userUUID = '';
      this.form.cartProducts = [];
    },
    clearResponseField: function() {
      this.resp.cart = [];
      this.resp.errorCode = 0;
      this.errorMsg = '';
    },
    getCart: function() {
      this.clearResponseField();
      const req = new GetRequest();
      req.setUseruuid(this.form.Useruuid);
      this.client.get(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let c = new Object();
          c.userUUID = resp.getCart().getUseruuid();
          c.cartProducts = resp.getCart().getCartproductsList();
          this.resp.cart.push(c);
          this.resp.errorCode = err.code;
        }
      });
    },
    setCart: function() {
      this.clearResponseField();
      const req = new SetRequest();
      const c = new Cart();
      c.setUseruuid(this.form.userUUID);
      var cartProducts = []
      this.form.cartProducts.forEach(function(v) {
        const cp = new CartProduct();
        cp.setProductuuid(v.productUUID);
        cp.setCount(v.count);
        cartProducts.push(cp)
      });
      c.setCartproductsList(cartProducts);
      req.setCart(c);
      this.client.set(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let c = new Object();
          c.Useruuid = resp.getUseruuid();
          this.resp.cart.push(c);
          this.resp.errorCode = err.code;
        }
      });
    },
    updateOrder: function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const c = new Cart();
      c.setUseruuid(this.form.userUUID);
      c.setCartproductsList(this.form.cartProducts);
      req.setCart(c);
      this.client.update(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
    deleteOrder: function() {
      this.clearResponseField();
      const req = new DeleteRequest();
      req.setUseruuid(this.form.userUUID);
      this.client.delete(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
  }
});
