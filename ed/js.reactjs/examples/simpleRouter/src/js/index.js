import React from 'react'
import { render } from 'react-dom'
import { browserHistory, Router, IndexRoute, Route, Link } from 'react-router'

// const App = ({ children, routes }) => {
class App extends React.Component {
    onNavigateHome() {
        browserHistory.push("/");
    }
    render() {
        return (
            <div>
                <aside>
                    <ul>
                        <li><Link to="/products" activeStyle={{color: "green"}}>Products</Link></li>
                        <li><Link to="/orders" activeClassName={"active"}>Orders</Link></li>
                        <li><Link to={"/order/" + (1 + 2)}>Order</Link></li>
                        <li>
                            <button onClick={this.onNavigateHome}>Home</button>
                        </li>
                    </ul>
                </aside>
                <main>
                    {this.props.children}
                </main>
            </div>
        )
    };
}

const Home = () => (
    <div className="Page">
        <h1>Home</h1>
    </div>
);

const Products = () => (
    <div className="Page">
        <h1>Products</h1>
    </div>
);

const Orders = () => (
    <div className="Page">
        <h1>Orders</h1>
    </div>
);

class Order extends React.Component {
    render() {
        return (
            <div>
                Order {this.props.params.id}
            </div>
        )
    };
}


render((
    <Router history={browserHistory}>
        <Route path="/" component={App}>
            <IndexRoute component={Home}/>
            <Route path="/products" component={Products} />
            <Route path="/orders" component={Orders} />
            <Route path="/order/:id" component={Order} />
        </Route>
    </Router>
), document.getElementById('root'));
