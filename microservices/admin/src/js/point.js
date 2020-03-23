const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, GetTotalAmountRequest, Point} = require('./protobuf/point_pb.js');

const {PointAPIClient} = require('./protobuf/point_grpc_web_pb.js');

export const point = new Vue({
  el: '#point',
  data: {
    endpoint: window.location.protocol + '//' + window.location.host + "/point",
    form: {
      uuid: '',
      userUUID: '',
      balance: null,
      description: '',
    },
    totalAmountForm: {
      userUUID: '',
    },
    resp: {
      point: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new PointAPIClient(this.endpoint);
  },
  methods: {
    clearForm: function() {
      this.form.uuid = '';
      this.form.userUUID = '';
      this.form.balance = null;
      this.form.description = '';
    },
    clearTotalAmountForm: function() {
      this.totalAmountForm.userUUID = '';
    },
    clearResponseField: function() {
      this.resp.point = [];
      this.resp.errorCode = 0;
      this.errorMsg = '';
    },
    getPoint: function() {
      this.clearResponseField();
      const req = new GetRequest();
      req.setUuid(this.form.uuid);
      this.client.get(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let p = new Object();
          p.uuid = resp.getPoint().getUuid();
          p.userUUID = resp.getPoint().getUseruuid();
          p.balance = resp.getPoint().getBalance();
          p.description = resp.getPoint().getDescription();
          p.createdAt = resp.getPoint().getCreatedat();
          p.updatedAt = resp.getPoint().getUpdatedat();
          p.deletedAt = resp.getPoint().getDeletedat();
          this.resp.point.push(p);
          this.resp.errorCode = err.code;
        }
      });
    },
    setPoint: function() {
      this.clearResponseField();
      const req = new SetRequest();
      const p = new Point();
      p.setUseruuid(this.form.userUUID);
      p.setBalance(this.form.balance);
      p.setDescription(this.form.description);

      req.setPoint(p);
      this.client.set(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let p = new Object();
          p.uuid = resp.getUuid();
          this.resp.point.push(p);
          this.resp.errorCode = err.code;
        }
      });
    },
    updatePoint: function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const p = new Point();
      p.setUuid(this.form.uuid);
      p.setUseruuid(this.form.userUUID);
      p.setBalance(this.form.balance);
      p.setDescription(this.form.description);
      req.setPoint(p);
      this.client.update(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
    deletePoint: function() {
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
    getTotalAmount: function() {
      this.clearResponseField();
      const req = new GetTotalAmountRequest();
      req.setUseruuid(this.totalAmountForm.userUUID);
      this.client.getTotalAmount(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let p = new Object();
          p.useruuid = resp.getUseruuid();
          p.totalAmount = resp.getTotalamount();
          this.resp.point.push(p);
          this.resp.errorCode = err.code;
        }
      });
    },
  }
});