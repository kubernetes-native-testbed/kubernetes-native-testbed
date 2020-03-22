const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, IsExistsRequest, IsExistsResponse, Product} = require('./protobuf/product_pb.js');
const {ImageUploadRequest, ImageUploadResponse, ImageDeleteRequest} = require('./protobuf/image_pb.js');

const {ProductAPIClient} = require('./protobuf/product_grpc_web_pb.js');
const {ImageAPIClient} = require('./protobuf/image_grpc_web_pb.js');

export const product = new Vue({
  el: '#product',
  data: {
    endpoint: window.location.protocol + '//' + window.location.host + "/product",
    form: {
      uuid: '',
      name: '',
      price: null,
      imageURLs: [],
      images: [],
    },
    resp: {
      product: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new ProductAPIClient(this.endpoint);
      this.imageClient = new ImageAPIClient(this.endpoint);
  },
  methods: {
    addImageURL: function() {
      this.form.imageURLs.push({value:''});
    },
    selectImage: function(e) {
      Object.keys(e.target.files).forEach((key) => {
        //console.log('image:', key, ':', e.target.files[key])
        var reader = new FileReader();
        reader.onload = () => {
          this.form.images.push(reader.result);
        }
        reader.readAsDataURL(e.target.files[key]);
      });
    },
    convertDataURIToBinary: function(dataURI) {
      var BASE64_MARKER = ';base64,';
      var base64Index = dataURI.indexOf(BASE64_MARKER) + BASE64_MARKER.length;
      var base64 = dataURI.substring(base64Index);
      var raw = window.atob(base64);
      var rawLength = raw.length;
      var array = new Uint8Array(new ArrayBuffer(rawLength));
      for(var i = 0; i < rawLength; i++) {
        array[i] = raw.charCodeAt(i);
      }
      return array;
    },
    uploadImageBlob: function(blob) {
      return new Promise((resolve, reject) => {
        let b = this.convertDataURIToBinary(blob)
        const imageReq = new ImageUploadRequest();
        imageReq.setImage(b);
        this.imageClient.upload(imageReq, {}, (err, resp) => {
          if (err) {
            this.resp.errorCode = err.code;
            this.resp.errorMsg = err.message;
          } else {
            resolve(resp);
          }
        });
      });
    },
    uploadImages: async function() {
      var promises = []
      this.form.images.forEach(async (v) => {
        promises.push(this.uploadImageBlob(v));
      });
      var resps = await Promise.all(promises)
      return resps
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
    setProduct: async function() {
      this.clearResponseField();
      const req = new SetRequest();
      const p = new Product();
      p.setName(this.form.name);
      p.setPrice(this.form.price);
      var urls = [];
      var resps = await this.uploadImages();
      resps.forEach(function(v) {
        //console.log("url:", v.getUrl());
        urls.push(v.getUrl());
      });
      console.log("urls:", urls);
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
    updateProduct: async function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const p = new Product();
      p.setUuid(this.form.uuid);
      p.setName(this.form.name);
      p.setPrice(this.form.price);
      var urls = [];
      var resps = await this.uploadImages();
      resps.forEach(function(v) {
        //console.log("url:", v.getUrl());
        urls.push(v.getUrl());
      });
      console.log("urls:", urls);
      p.setImageurlsList(urls);
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
    isExistsProduct: function() {
      this.clearResponseField();
      const req = new IsExistsRequest();
      req.setUuid(this.form.uuid);
      this.client.isExists(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let p = new Object();
          p.productUUID = this.form.uuid;
          p.isExists = resp.getIsexists();
          this.resp.product.push(p);
          this.resp.errorCode = err.code;
        }
      });
    },
  }
});
