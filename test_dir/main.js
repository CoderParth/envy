// This file is only for test purposes. 
//

const express = require('express');
const mongoose = require('mongoose');
const AWS = require('aws-sdk');
const sgMail = require('@sendgrid/mail');

const app = express();

// Connect to MongoDB
mongoose.connect(process.env.DB_URL, { useNewUrlParser: true, useUnifiedTopology: true });

// Configure AWS
AWS.config.update({
  accessKeyId: process.env.AWS_ACCESS_KEY_ID,
  secretAccessKey: process.env.AWS_ACCESS_KEY_SECRET,
  region: 'us-west-2'
});

// Configure SendGrid
sgMail.setApiKey(process.env.SENDGRID_API_KEY);

app.get('/', (req, res) => {
  res.send('Hello, World!');
});

app.listen(3000, () => {
  console.log('Server is running on port 3000');
});
