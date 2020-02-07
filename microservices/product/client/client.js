const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, Product} = require('./product_pb.js');
const {ProductAPIClient} = require('./product_grpc_web_pb.js');
const {Empty} = require('google-protobuf/google/protobuf/empty_pb.js');

const defaultEndpoint = 'https://testbed.amsy.dev/product';

const product = new Vue({
  el: '#product',
  data: {
    endpoint: defaultEndpoint,
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
          console.log(err.code);
          console.log(err.message);
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          console.log(resp.getProduct().getUuid());
          console.log(resp.getProduct().getName());
          console.log(resp.getProduct().getPrice());
          console.log(resp.getProduct().getImageurlsList());
          console.log(resp.getProduct().getCreatedat());
          console.log(resp.getProduct().getUpdatedat());
          console.log(resp.getProduct().getDeletedat());
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
          console.log(err.code);
          console.log(err.message);
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          console.log(resp.getUuid());
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
          console.log(err.code);
          console.log(err.message);
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
          console.log(err.code);
          console.log(err.message);
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
  }
});
