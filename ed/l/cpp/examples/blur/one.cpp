/*
@example:
docker run -it --rm -v $PWD:/gh -w /gh spmallick/opencv-docker:opencv sh -c '
  f=ed/l/cpp/examples/blur/one.cpp
  g++ -w $f -o x `pkg-config --cflags --libs opencv`
  ./x /gh/ed/l/python/examples/blur/z.png
'
*/

#include <opencv2/core/core.hpp> 
#include <opencv2/highgui/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>
#include <stdio.h>
#include <iostream>

using namespace cv;
using namespace std;

cv::Rect gr(int h, int w)
{
    int top    = int(0.14 * h);
    int bottom = int(0.25 * h);
    int left   = int(0.60 * w);
    int right  = int(0.71 * w);

    return cv::Rect(left, top, right - left, bottom - top);
}

int b(string f)
{
    cv::Mat image = imread(f, CV_LOAD_IMAGE_UNCHANGED);
    if (!image.data)
    {
        cout << "Could not open or find the image.\n";
        return -1;
    }

    cv::Size s = image.size();
    int h = s.height;
    int w = s.width;

    cv::Rect rc = gr(h, w);
    cv::Mat r = image(rc);
    blur(r, r, Size(50, 50));
    image(rc) = r;

    // namedWindow("blur", CV_WINDOW_AUTOSIZE);
    // imshow("blur", image);
    // waitKey(0);
    imwrite(f+".r.png", image);
}

int main(int argc, char *argv[])
{
    string f = argv[1];
    b(f);

    return 0;
}
