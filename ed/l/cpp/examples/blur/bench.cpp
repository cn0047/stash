/*
@example:
docker run -it --rm -v $PWD:/gh -w /gh spmallick/opencv-docker:opencv sh -c '
  f=ed/l/cpp/examples/blur/bench.cpp
  g++ -w $f -o x `pkg-config --cflags --libs opencv`
  ./x
'
*/

#include <opencv2/core/core.hpp>
#include <opencv2/highgui/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>
#include <stdio.h>
#include <iostream>

int b(std::string f)
{
    cv::Mat img = cv::imread(f, CV_LOAD_IMAGE_UNCHANGED);

    int top = 286, bottom = 308, left = 111, right = 144;
    cv::Rect r = cv::Rect(left, top, right - left, bottom - top);
    cv::Mat rec = img(r);
    cv::blur(rec, rec, cv::Size(3, 9));
    img(r) = rec;

    cv::imwrite(f+".r.png", img);
}

int main(int argc, char *argv[])
{
    b("/gh/ed/l/python/examples/blur/a.png");
    return 0;
}
