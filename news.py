#!/usr/bin/env python3
# -*- coding: UTF-8 -*-

import sys
import argparse
import time
from webbrowser import open_new_tab
from math import trunc

try:
    import tkinter
    import tkinter.font
except ImportError:
    sys.stderr.write('your python do not support tkinter')
    sys.exit(1)

def ago(when):
    now = time.time()
    sec = trunc(now - when)
    min, sec = divmod(sec, 60)

    if min > 0:
        return '{}m {}s'.format(min, sec)

    return '{}s'.format(sec)

def make_card(name, text, when, url):
    window = tkinter.Tk()

    def update():
        window.wm_title('{} - {} 전'.format(name, ago(when)))
        window.after(1000, update)

    update()

    font = tkinter.font.Font(window, ('맑은 고딕', 24, ''))

    text_lab = tkinter.Label(window, text=text, font=font)
    text_lab.pack(expand=True, padx=10, pady=10)

    def open_url(event):
        open_new_tab(url)

    text_lab.bind("<Button-1>", open_url)

    window.wm_withdraw()
    window.update()

    width = max(200, window.winfo_width())
    height = window.winfo_height()
    window.wm_geometry('{}x{}'.format(width, height))
    window.wm_minsize(width, height)
    window.wm_maxsize(width, height)

    window.bell()
    window.wm_deiconify()
    window.mainloop()

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('name')
    parser.add_argument('text')
    parser.add_argument('when', type=int)
    parser.add_argument('url')

    args = parser.parse_args()
    make_card(args.name, args.text, args.when, args.url)

if __name__ == '__main__':
    main()
