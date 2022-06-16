import React from 'react';
import Navbar from '../../../component/navbar/navbar';
import Footer from '../../../component/footer/footer';
import {MenuFakultasTeknik} from './menuFakultasTeknik';
import Button from '../../../component/button/button';
import {Link} from 'react-router-dom';
import './fakultasTeknik.css';

const FakultasTeknik = () => {
	return (
		<div>
			<Navbar />
			<div class="MainSection">
				<div className="heading">
					<h1>List Program Studi Fakultas Teknik</h1>
					<p>Simak Program Studi yang sesuai dengan passion kalian</p>
				</div>
				<div className="Main">
					<ul className="main-menu-list">
						{MenuFakultasTeknik.map((item, index) => (
							<li key={index} className="fakultas-list">
								<h1 className="header">RightWay</h1>
								<div className="link-text">
									<Link className={item.cName} to={item.url}>
										{item.title}
									</Link>
									<Link
										to={item.url}
										className="selengkapnya"
									>
										Lihat selengkapnya
									</Link>
								</div>
							</li>
						))}
					</ul>
				</div>
				<div className="btn-see-more">
					<Link to="/sign-up">
						<Button
							buttonStyle="btn--primary"
							children="Lihat fakultas lainnya"
						/>
					</Link>
				</div>
			</div>
			<Footer />
		</div>
	);
};

export default FakultasTeknik;
