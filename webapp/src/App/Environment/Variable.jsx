import React, {Component} from 'react';
import Input from '../Input';
import './variable.css';

export default class Entry extends Component {
  state = this.props.value;

  handleValidName = (name) => {
    this.setState(prevState => {
      const value = Object.assign(prevState, {name});
      this.props.onValid(value);
      return value;
    });
  };

  handleValidValue = (value) => {
    this.setState({value});
  };

  render() {
    const {onInvalid} = this.props;
    return (
      <li className='list-group-item'>
        <div className='row'>
          <div className='col-xl-4'>
            <Input
              id={`${this.state.id}-name`}
              onInvalid={onInvalid}
              onValid={this.handleValidName}
              type={'text'}
              validate={entryName => entryName.length > 0 ? [] : [{message: "name required"}]}
              value={this.state.name}
            />
          </div>
          <div className='col-xl-8'>
            <Input
              id={`${this.state.id}-value`}
              onInvalid={onInvalid}
              onValid={this.handleValidValue}
              type={'text'}
              validate={() => []}
              value={this.state.value}
            />
          </div>
        </div>
      </li>
    );
  }
}