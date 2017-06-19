@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row">
        <div class="col-md-10 col-md-offset-1">
            <div class="panel panel-default">
                <div class="panel-heading">Welcome</div>

                <div class="panel-body">
                    Your Application's Landing Page.
                    You can try something from:
                    | <a href="{{ url('/get/200') }}">get 200</a>
                    | <a href="{{ url('/home') }}">home</a>
                    | <a href="{{ url('/di') }}">di</a>
                    | <a href="{{ url('/db') }}">db</a>
                    | <a href="{{ url('/auth') }}">auth</a>
                    | <a href="{{ url('/conf') }}">conf</a>
                    | <a href="{{ url('/flash') }}">flash</a>
                    | <a href="{{ url('/validation') }}">validation</a>
                    | <a href="{{ url('/localization') }}">localization</a>
                    @if (Auth::check())
                        <hr>
                        For logged users:
                        | <a href="{{ url('/userFromDb') }}">userFromDb</a>
                    @endif
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
