import AWS from 'aws-sdk';
import React, { Component } from 'react';

import config from './../../config/main';

class Cell1 extends Component {

  constructor(props) {
    super(props);
    this.state = {img: '', data: props};
  }

  componentDidMount() {
    AWS.config.update(config.aws.mail);
    let s3 = new AWS.S3();
    let p = this.state.data.pictures[0];
    let args = {
      Bucket: config.aws.bucket,
      Key: this.state.data.bucket_folder + '/public/' + p.file_name.replace('.jpg', '_thumbnail.jpg')
    };
    let that = this;
    s3.getSignedUrl('getObject', args, (err, url) => {
      that.setState({img: url});
    });
  }

  render() {
    return (
      <div className="ssc2">
        <div className="ct">{this.props.user_id}</div>
        <img className="img" src={this.state.img} alt="" />
        <div className="cb">
          {this.props.favourited > 0 ? '⭐' : ''}
          {this.props.flamed > 0 ? '🔥' : ''}
          {this.props.visited > 0 ? '👁' : ''}
          {this.props.profile_meta_data.name}
        </div>
      </div>
    );
  }

}

export default Cell1;
