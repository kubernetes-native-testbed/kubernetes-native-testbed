const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, PaymentInfo} = require('./protobuf/payment_info_pb.js');

const {PaymentInfoAPIClient} = require('./protobuf/payment_info_grpc_web_pb.js');

export const paymentInfo = new Vue({
  el: '#paymentInfo',
  data: {
    endpoint: window.location.protocol + '//' + window.location.host + "/paymentinfo",
    form: {
      uuid: '',
      userUUID: '',
      name: '',
      cardNumber: '',
      securityCode: '',
      expirationDate: '',
    },
    resp: {
      paymentInfo: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new PaymentInfoAPIClient(this.endpoint);
  },
  methods: {
    clearForm: function() {
      this.form.uuid = '';
      this.form.userUUID = '';
      this.form.name = '';
      this.form.cardNumber = '';
      this.form.securityCode = '';
      this.form.expirationDate = '';
    },
    clearResponseField: function() {
      this.resp.paymentInfo = [];
      this.resp.errorCode = 0;
      this.errorMsg = '';
    },
    getPaymentInfo: function() {
      this.clearResponseField();
      const req = new GetRequest();
      req.setUuid(this.form.uuid);
      this.client.get(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let pi = new Object();
          pi.uuid = resp.getPaymentInfo().getUuid();
          pi.userUUID = resp.getPaymentInfo().getUseruuid();
          pi.name = resp.getPaymentInfo().getName();
          pi.cardNumber = resp.getPaymentInfo().getCardnumber();
          pi.securityCode = resp.getPaymentInfo().getSecuritycode();
          pi.expirationDate = resp.getPaymentInfo().getExpirationdate();
          pi.createdAt = resp.getPaymentInfo().getCreatedat();
          pi.updatedAt = resp.getPaymentInfo().getUpdatedat();
          pi.deletedAt = resp.getPaymentInfo().getDeletedat();
          this.resp.paymentInfo.push(pi);
          this.resp.errorCode = err.code;
        }
      });
    },
    setPaymentInfo: function() {
      this.clearResponseField();
      const req = new SetRequest();
      const pi = new PaymentInfo();
      pi.setUseruuid(this.form.userUUID);
      pi.setName(this.form.name);
      pi.setCardnumber(this.form.cardNumber);
      pi.setSecuritycode(this.form.securityCode);
      pi.setExpirationdate(this.form.expirationDate);

      req.setPaymentinfo(pi);
      this.client.set(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let pi = new Object();
          pi.uuid = resp.getUuid();
          this.resp.paymentInfo.push(pi);
          this.resp.errorCode = err.code;
        }
      });
    },
    updatePaymentInfo: function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const pi = new PaymentInfo();
      pi.setUseruuid(this.form.userUUID);
      pi.setName(this.form.name);
      pi.setCardnumber(this.form.cardNumber);
      pi.setSecuritycode(this.form.securityCode);
      pi.setExpirationdate(this.form.expirationDate);
      req.setPaymentinfo(pi);
      this.client.update(req, {}, (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
    deletePaymentInfo: function() {
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