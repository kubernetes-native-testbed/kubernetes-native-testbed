const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, Product} = require('./protobuf/product_pb.js');
const {ProductAPIClient} = require('./protobuf/product_grpc_web_pb.js');
const {Empty} = require('google-protobuf/google/protobuf/empty_pb.js');

const defaultEndpoint = 'https://34.84.105.184.nip.io';

const product = new Vue({
  el: '#product',
  data: {
    endpoint: defaultEndpoint + "/product",
    form: {
      uuid: '',
      name: '',
      price: 0,
      imageURLs: [],
    },
    resp: {
      product: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new ProductAPIClient(this.endpoint);
  },
  methods: {
    addImageURL: function() {
      this.form.imageURLs.push({value:''});
    },
    clearForm: function() {
      this.form.uuid = '';
      this.form.name = '';
      this.form.price = null;
      this.form.imageURLs = [];
    },
    clearResponseField: function() {
      this.resp.product = [];
      this.resp.errorCode = 0;
      this.errorMsg = '';
    },
    getProduct: function() {
      this.clearResponseField();
      const req = new GetRequest();
      req.setUuid(this.form.uuid);
      this.client.get(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let p = new Object();
          p.uuid = resp.getProduct().getUuid();
          p.price = resp.getProduct().getPrice();
          p.imageURLs = resp.getProduct().getImageurlsList();
          p.createdAt = resp.getProduct().getCreatedat();
          p.updatedAt = resp.getProduct().getUpdatedat();
          p.deletedAt = resp.getProduct().getDeletedat();
          this.resp.product.push(p);
          this.resp.errorCode = err.code;
        }
      });
    },
    setProduct: function() {
      this.clearResponseField();
      const req = new SetRequest();
      const p = new Product();
      p.setName(this.form.name);
      p.setPrice(this.form.price);
      urls = []
      this.form.imageURLs.forEach(function(v) {
        urls.push(v.value)
      });
      p.setImageurlsList(urls);
      req.setProduct(p);
      this.client.set(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let p = new Object();
          p.uuid = resp.getUuid();
          this.resp.product.push(p);
          this.resp.errorCode = err.code;
        }
      });
    },
    updateProduct: function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const p = new Product();
      p.setUuid(this.form.uuid);
      p.setName(this.form.name);
      p.setPrice(this.form.price);
      p.setImageurlsList(this.form.imageURLs);
      req.setProduct(p);
      this.client.update(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
    deleteProduct: function() {
      this.clearResponseField();
      const req = new DeleteRequest();
      req.setUuid(this.form.uuid);
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

const user = new Vue({
  el: '#user',
  data: {
    endpoint: defaultEndpoint + "/user",
    form: {
      uuid: '',
      username: '',
      firstName: '',
      lastName: '',
      age: 0,
      passwordHash: '',
      addresses: [],
    },
    resp: {
      user: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new UserAPIClient(this.endpoint);
  },
  methods: {
    addAddresses: function() {
      this.form.addresses.push({value:''});
    },
    clearForm: function() {
      this.form.uuid = '',
      this.form.username = '',
      this.form.firstName = '',
      this.form.lastName = '',
      this.form.age = null,
      this.form.passwordHash = '',
      this.form.addresses = [];
    },
    clearResponseField: function() {
      this.resp.user = [];
      this.resp.errorCode = 0;
      this.errorMsg = '';
    },
    getUser: function() {
      this.clearResponseField();
      const req = new GetRequest();
      req.setUuid(this.form.uuid);
      this.client.get(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let u = new Object();
          u.uuid = resp.getUser().getUuid();
          u.username = resp.getUser().getUsername();
          u.firstName = resp.getUser().getFirstName();
          u.lastName = resp.getUser().getLastName();
          u.age = resp.getUser().getAge();
          u.passwordHash = resp.getUser().getPasswordHash();
          u.addresses = resp.getUser().getAddressList();
          u.createdAt = resp.getUser().getCreatedat();
          u.updatedAt = resp.getUser().getUpdatedat();
          u.deletedAt = resp.getUser().getDeletedat();
          this.resp.user.push(u);
          this.resp.errorCode = err.code;
        }
      });
    },
    setUser: function() {
      this.clearResponseField();
      const req = new SetRequest();
      const u = new User();
      u.setUsername(this.form.username);
      u.setFirstName(this.form.firstName);
      u.setLastName(this.form.lastName);
      u.setAge(this.form.age);
      u.setPasswordHash(this.form.passwordHash);

      addresses = []
      this.form.addresses.forEach(function(v) {
        addresses.push(v.value)
      });
      u.setAddressesList(addresses);
      req.setUser(u);
      this.client.set(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let u = new Object();
          u.uuid = resp.getUuid();
          this.resp.user.push(u);
          this.resp.errorCode = err.code;
        }
      });
    },
    updateUser: function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const u = new User();
      u.setUsername(this.form.username);
      u.setFirstName(this.form.firstName);
      u.setLastName(this.form.lastName);
      u.setAge(this.form.age);
      u.setPasswordHash(this.form.passwordHash);
      u.setAddressesList(this.form.imageURLs);
      req.setUser(u);
      this.client.update(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
    deleteUser: function() {
      this.clearResponseField();
      const req = new DeleteRequest();
      req.setUuid(this.form.uuid);
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
