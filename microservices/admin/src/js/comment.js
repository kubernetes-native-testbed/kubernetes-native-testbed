const {GetRequest, GetResponse, SetRequest, SetResponse, UpdateRequest, DeleteRequest, Comment} = require('./protobuf/comment_pb.js');

const {CommentAPIClient} = require('./protobuf/comment_grpc_web_pb.js');

const {GetTokenMetadata} = require('./cookie.js');

export const comment = new Vue({
  el: '#comment',
  data: {
    endpoint: window.location.protocol + '//' + window.location.host + "/comment",
    form: {
      uuid: '',
      userUUID: '',
      parentCommentUUID: '',
      message: '',
    },
    resp: {
      comment: [],
      errorCode: 0,
      errorMsg: '',
    }
  },
  created: function() {
      this.client = new CommentAPIClient(this.endpoint);
  },
  methods: {
    clearForm: function() {
      this.form.uuid = '';
      this.form.userUUID = '';
      this.form.parentCommentUUID = '';
      this.form.message = '';
    },
    clearResponseField: function() {
      this.resp.comment = [];
      this.resp.errorCode = 0;
      this.errorMsg = '';
    },
    getComment: function() {
      this.clearResponseField();
      const req = new GetRequest();
      req.setUuid(this.form.uuid);
      this.client.get(req, GetTokenMetadata(), (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let o = new Object();
          o.uuid = resp.getComment().getUuid();
          o.userUUID = resp.getComment().getUseruuid();
          o.parentCommentUUID = resp.getComment().getParentcommentuuid();
          o.message = resp.getComment().getMessage();
          o.createdAt = resp.getComment().getCreatedat();
          o.updatedAt = resp.getComment().getUpdatedat();
          o.deletedAt = resp.getComment().getDeletedat();
          this.resp.comment.push(o);
          this.resp.errorCode = err.code;
        }
      });
    },
    setComment: function() {
      this.clearResponseField();
      const req = new SetRequest();
      const c = new Comment();
      c.setUseruuid(this.form.userUUID);
      c.setParentcommentuuid(this.form.parentCommentUUID);
      c.setMessage(this.form.message);
      req.setComment(c);
      this.client.set(req, GetTokenMetadata(), (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          let o = new Object();
          o.uuid = resp.getUuid();
          this.resp.comment.push(o);
          this.resp.errorCode = err.code;
        }
      });
    },
    updateComment: function() {
      this.clearResponseField();
      const req = new UpdateRequest();
      const c = new Comment();
      c.setUuid(this.form.uuid);
      c.setUseruuid(this.form.userUUID);
      c.setParentcommentuuid(this.form.parentCommentUUID);
      c.setMessage(this.form.message);
      req.setComment(c);
      this.client.update(req, GetTokenMetadata(), (err, resp) => {
        if (err) {
          this.resp.errorCode = err.code;
          this.resp.errorMsg = err.message;
        } else {
          this.resp.errorCode = err.code;
        }
      });
    },
    deleteComment: function() {
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
