import React, {Component} from 'react';
import AWS from 'aws-sdk';

import { MainConfig } from './../Config/Main';

class Cell1 extends Component {

  constructor(props) {
    super(props);
    this.state = {img: '', data: props};
  }

  componentDidMount() {
    if (typeof this.state.data.pictures[0] === 'undefined') {
      return;
    }
    let p = this.state.data.pictures[0];
    if (p.type_id !== 201201) {
      return;
    }
    AWS.config.update(MainConfig.aws.mail);
    let s3 = new AWS.S3();
    let args = {
      Bucket: MainConfig.aws.bucket,
      Key: this.state.data.bucket_folder + '/public/' + p.file_name.replace('.jpg', '_thumbnail.jpg')
    };
    let that = this;
    s3.getSignedUrl('getObject', args, (err, url) => {
      that.setState({img: url});
    });
  }

  render() {
    let img = '';
    if (this.state.img === '') {
      img = <div className="noImg">No image</div>;
    } else {
      img = <img className="img" src={this.state.img} alt="" />;
    }
    return (
      <div className="ssc2">
        <div className="ct">{this.props.user_id}</div>
        {img}
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
