import React from 'react';
import './App.css';
import { Header } from './compornets/Header';
import { TodoList } from './compornets/TodoList';

function App() {
  return (
    <div className="App">
      <Header />
      <TodoList todos={[
        {title: "Do dishes", description: "Do dishes", isCompleted: true},
        {title: "Do dishes", description: "Do dishes", isCompleted: true},
      ]} />
    </div>
  );
}

export default App;
