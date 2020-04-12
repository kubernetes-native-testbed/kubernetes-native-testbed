const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, Rate} = require('./protobuf/rate_pb.js');

const {RateAPIClient} = require('./protobuf/rate_grpc_web_pb.js');

const {GetTokenMetadata} = require('./cookie.js');

export const rate = new Vue({
  el: '#rate',
  data: {
    endpoint: window.location.protocol + '//' + window.location.host + "/rate",
    form: {
      uuid: '',
      userUUID: '',
      commentUUID: '',
      productUUID: '',
      rating: null,
    },
    resp: {
      rate: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new RateAPIClient(this.endpoint);
  },
  methods: {
    clearForm: function() {
      this.form.uuid = '';
      this.form.userUUID = '';
      this.form.commentUUID = '';
      this.form.productUUID = '';
      this.form.rating = null;
    },
    clearResponseField: function() {
      this.resp.rate = [];
      this.resp.errorCode = 0;
      this.errorMsg = '';
    },
    getRate: function() {
      this.clearResponseField();
      const req = new GetRequest();
      req.setUuid(this.form.uuid);
      this.client.get(req, GetTokenMetadata(), (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let o = new Object();
          o.uuid = resp.getRate().getUuid();
          o.userUUID = resp.getRate().getUseruuid();
          o.commentUUID = resp.getRate().getCommentuuid();
          o.productUUID = resp.getRate().getProductuuid();
          o.rating = resp.getRate().getRating();
          o.createdAt = resp.getRate().getCreatedat();
          o.updatedAt = resp.getRate().getUpdatedat();
          o.deletedAt = resp.getRate().getDeletedat();
          this.resp.rate.push(o);
          this.resp.errorCode = err.code;
        }
      });
    },
    setRate: function() {
      this.clearResponseField();
      const req = new SetRequest();
      const r = new Rate();
      r.setUseruuid(this.form.userUUID);
      r.setCommentuuid(this.form.commentUUID);
      r.setProductuuid(this.form.productUUID);
      r.setRating(this.form.rating);
      req.setRate(r);
      this.client.set(req, GetTokenMetadata(), (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let o = new Object();
          o.uuid = resp.getUuid();
          this.resp.rate.push(o);
          this.resp.errorCode = err.code;
        }
      });
    },
    updateRate: function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const r = new Rate();
      r.setUuid(this.form.uuid);
      r.setUseruuid(this.form.userUUID);
      r.setCommentuuid(this.form.commentUUID);
      r.setProductuuid(this.form.productUUID);
      r.setRating(this.form.rating);
      req.setRate(r);
      this.client.update(req, GetTokenMetadata(), (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
    deleteRate: function() {
      this.clearResponseField();
      const req = new DeleteRequest();
      req.setUuid(this.form.uuid);
      this.client.delete(req, GetTokenMetadata(), (err, resp) => {
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
