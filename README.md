## Julia Set Generator Written In Golang 

This is a simple (naive) Julia Set generator written in Golang. The project utilizes concurrent workers to speed up processing speed. This project 
creates a set of png images which can be used by Ffmpeg to create an mp4 video showing an animated julia set.

### Dependencies 
FFMPEG - video processing tool which can produce a video from a set of images
github.com/muesli/gamut - for generating color pallets. 

### Needed Improvements 
+ In memory frame storage 
    + Currently, each frame of the desired animation will be stored as a png file on your computers file system. This uses a large amount of 
    hard drive space and likely slows the entire process down somewhat. Ideally, frames should be stored in memory until they are passed to ffmpeg. Currently, 
    they are stored within the filesystem for simplicity, as it allows you to just easily run a ffmpeg command to produce a mp4 video
+ Improved efficiency
    + I'm sure there are a few ways to speed this project up computationally, aside from just throwing more CPU cores at it.
+ Automatic FFMPEG video creation 
    + Having this tool automatically create a video would be cool, as opposed to having to manually run an FFMPEG command. 

### Some Good Starting values 
Different julia sets can be created by using different values for the constant real and constant imaginary values, here are some good values to produce a few julia sets that are interesting to look at 

| cReal | cImaginary | Increment | endRange |
|-------|------------|-----------|----------|
| -0.8  | 0.156      |  0.00001  | 0.0028   |
| -0.4  | 0.6        |  0.0001   | .05      |
| 0.280 | 0.01       |  0.00001   | 0.005     | 

additional values can be found on various sites, such as wikipedia. 

### How to create a video of a julia set

Using this repository will create a set of png's,
in essence, video frames. to produce an actual mp4
you can use the following FFMPEG commands.
Ensure that you run these commands within the repository. These commands could likely be improved to provide greater image quality and better compression.

Windows: `ffmpeg -framerate 30 -i img%04d.png -c:v libx264 -pix_fmt yuv420p out.mp4`

MacOS: `ffmpeg -framerate 30 -pattern_type glob -i '*.png' -c:v libx264 -pix_fmt yuv420p out.mp4`

### What is a julia set?

According to [Britannica](https://www.britannica.com/science/Julia-set),

```text
In general terms, a Julia set is the boundary between points in the complex number
plane or the Riemann sphere (the complex number plane plus the point at infinity)
that diverge to infinity and those that remain finite under repeated iteration of
some mapping (function). 
```

Julia set's are an interesting view into fractals and how mathematics and art can intersect. Even without an in depth 
appreciation for the mathematics at play, almost everyone can appreciate the resulting images / videos. Here are some examples, 


![example 1](./example-1.png)
![example 2](./example-2.png)
![example 3](./example-3.png)
