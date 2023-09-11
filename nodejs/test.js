const express = require('express');
const chai = require('chai');
const chaiHttp = require('chai-http');
const { expect } = chai;

const server = require('./server.js'); 

chai.use(chaiHttp);

describe('Task API', () => {
  before(() => {
    
  });

  after(() => {
   
  });

  beforeEach(() => {
   
  });

  it('should return a list of tasks', (done) => {
    chai
      .request(server)
      .get('/tasks')
      .end((err, res) => {
        expect(res).to.have.status(200);
        expect(res.body).to.be.an('array');
        done();
      });
  });

  it('should return "OK" for health check', (done) => {
    chai
      .request(server)
      .get('/healthz')
      .end((err, res) => {
        expect(res).to.have.status(200);
        expect(res.body).to.equal('OK');
        done();
      });
  });

  it('should create a new task', (done) => {
    const task = {
      title: 'Test Task',
      description: 'This is a test task',
    };

    chai
      .request(server)
      .post('/tasks')
      .send(task)
      .end((err, res) => {
        expect(res).to.have.status(201);
        expect(res.body).to.have.property('title', task.title);
        expect(res.body).to.have.property('description', task.description);

        done();
      });
  });
});