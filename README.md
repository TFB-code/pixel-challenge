# pixel-challenge

The Challenge

Given a selection of images, and choosing a particular reference one, find the 3
“closest” images to that reference image.
How “close” an image is determined by the number of pixels that match precisely
between the two images - both in value and location.
% Closeness = (Number of Pixels Matching) / (Number of Pixels Total) * 100
The output should be a list of images, and their to the reference either as full
percentage (97%) or decimal (0.97)

The Images

Images are purely random noise. They are 1024 pixels wide
and 1024 pixels high.
Each pixel is encoded as 3 bytes of data, representing the
Red, Green and Blue Channels.
For example: the first byte could be represented as [255, 255,
255] which would be a white pixel. Another pixel could be [0, 0,
0] which would be a black pixel.

Levels

As with all games, there are levels!

Bronze - 10 Images in dataset (30 megabytes)
https://pixels-data.s3.eu-west-2.amazonaws.com/Bronze.zip

Silver - 100 Images in dataset (300 megabytes)
https://pixels-data.s3.eu-west-2.amazonaws.com/Silver.zip

Gold - 1000 Images in dataset (3000 megabytes)
https://pixels-data.s3.eu-west-2.amazonaws.com/Gold.zip
