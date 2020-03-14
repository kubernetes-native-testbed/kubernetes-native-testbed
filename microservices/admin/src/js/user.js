export const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, User, Address} = require('./protobuf/user_pb.js');

export const {UserAPIClient} = require('./protobuf/user_grpc_web_pb.js');

export const user = new Vue({
  el: '#user',
  data: {
    endpoint: window.location.protocol + '//' + window.location.host + "/user",
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
    addAddress: function() {
      this.form.addresses.push({zipCode:'', country: '', state: '', city: '', addressLine: '', disabled: null});
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
          u.firstName = resp.getUser().getFirstname();
          u.lastName = resp.getUser().getLastname();
          u.age = resp.getUser().getAge();
          u.passwordHash = resp.getUser().getPasswordhash();
          u.addresses = resp.getUser().getAddressesList();
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
      u.setFirstname(this.form.firstName);
      u.setLastname(this.form.lastName);
      u.setAge(this.form.age);
      u.setPasswordhash(this.form.passwordHash);

      var addresses = []
      this.form.addresses.forEach(function(v) {
        const a = new Address();
        a.setZipcode(v.zipCode)
        a.setCountry(v.country)
        a.setState(v.state)
        a.setCity(v.city)
        a.setAddressline(v.addressLine)
        a.setDisabled(v.disabled)
        addresses.push(a)
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
      u.setFirstname(this.form.firstName);
      u.setLastname(this.form.lastName);
      u.setAge(this.form.age);
      u.setPasswordhash(this.form.passwordHash);
      u.setAddressesList(this.form.addresses);
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
