import React, { Component } from 'react';
import Form from './Form';
import List from './List';
import axios from 'axios';
export default class App extends Component {
  constructor(props){
    super(props);
  	this.state = {
	    todo: []
	};
    axios.get("http://localhost:8888/").then(res=>{
  		console.log(res.data);
  		this.setState({todo: res.data});
  	});
    this.handleAdd = this.handleAdd.bind(this);
    this.handleRemove = this.handleRemove.bind(this);
  }
  // データ保存
  handleAdd(e){
    e.preventDefault();
    //データベースに保存
    const data = {'name': e.target.title.value};
    axios.post('http://localhost:8888/todo/post',data).then(res=>{
    	var id = res.data.ID;
    	var name = res.data.Name
    	 // フォームから受け取ったデータをオブジェクトに挿入して、stateのtodo配列に追加
    	this.state.todo.push({ID: id, Name: name});
    	this.setState({todo: this.state.todo});
    });
    // inputのvalueを空に
    e.target.title.value = '';
  }

  // データ削除
  handleRemove(i,id){
    // todo配列からi番目から1つ目のデータを除外
    console.log(id);
    this.state.todo.splice(i,1);
    // setStateでtodo配列を上書き
    this.setState({todo: this.state.todo});
    // データベースから削除する
    axios.get('http://localhost:8888/todo/delete/' + id);
  }
  render() {
    return (
      <div className="siimple-box siimple--bg-dark">
        <h1 className="siimple-box-title siimple--color-white">React Todo App</h1>
        <Form handleAdd={this.handleAdd}/>
        <div className="siimple-rule"></div>
        <List todos={this.state.todo} handleRemove={this.handleRemove}/>
      </div>
    );
  }
}
