const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, DeliveryStatus} = require('./protobuf/delivery_status_pb.js');

const {DeliveryStatusAPIClient} = require('./protobuf/delivery_status_grpc_web_pb.js');

export const order = new Vue({
  el: '#delivery-status',
  data: {
    endpoint: window.location.protocol + '//' + window.location.host + "/delivery-status",
    form: {
      orderUUID: '',
      status: '',
      inquiryNumber: '',
    },
    resp: {
      deliveryStatus: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new DeliveryStatusAPIClient(this.endpoint);
  },
  methods: {
    clearForm: function() {
      this.form.orderUUID = '';
      this.form.status = '';
      this.form.inquiryNumber = '';
    },
    clearResponseField: function() {
      this.resp.deliveryStatus = [];
      this.resp.errorCode = 0;
      this.errorMsg = '';
    },
    getDeliveryStatus: function() {
      this.clearResponseField();
      const req = new GetRequest();
      req.setOrderuuid(this.form.orderUUID);
      this.client.get(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let s = new Object();
          s.orderUUID = resp.getDeliveryStatus().getOrderuuid();
          s.status = resp.getDeliveryStatus().getStatus();
          s.inquiryNumber = resp.getDeliveryStatus().getInquirynumber();
          s.createdAt = resp.getDeliveryStatus().getCreatedat();
          s.updatedAt = resp.getDeliveryStatus().getUpdatedat();
          s.deletedAt = resp.getDeliveryStatus().getDeletedat();
          this.resp.deliveryStatus.push(s);
          this.resp.errorCode = err.code;
        }
      });
    },
    setDeliveryStatus: function() {
      this.clearResponseField();
      const req = new SetRequest();
      const s = new DeliveryStatus();
      s.setOrderuuid(this.form.orderUUID);
      s.setStatus(this.form.status);
      s.setInquirynumber(this.form.inquiryNumber);
      req.setDeliverystatus(s);
      this.client.set(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let s = new Object();
          s.orderUUID = resp.getOrderuuid();
          this.resp.deliveryStatus.push(s);
          this.resp.errorCode = err.code;
        }
      });
    },
    updateDeliveryStatus: function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const s = new DeliveryStatus();
      s.setOrderuuid(this.form.orderUUID);
      s.setStatus(this.form.status);
      s.setInquirynumber(this.form.inquiryNumber);
      req.setDeliverystatus(s);
      this.client.update(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
    deleteDeliveryStatus: function() {
      this.clearResponseField();
      const req = new DeleteRequest();
      req.setOrderuuid(this.form.orderUUID);
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
