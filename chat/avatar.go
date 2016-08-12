package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL はAvatarインスタンスがアバターのURLを返せないときに投げられる
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatar はユーザのプロフィールを表す型
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返す
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars は3通りのアバター取得方法を順次試していくための型
type TryAvatars []Avatar

// AuthAvatar は、各種SNSの認証のための型
type AuthAvatar struct{}

// UseAuthAvatar は、アバター取得に各種SNSを用いることを表す型
var UseAuthAvatar AuthAvatar

// GetAvatarURL は、3通りの方法でユーザのアバターのURLを取得する
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GetAvatarURL は、3通りの方法でユーザのアバターのURLを取得する
func (a AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar は、Gravatarの認証のための型
type GravatarAvatar struct{}

// UseGravatar は、アバター取得にGravatarを用いることを表す型
var UseGravatar GravatarAvatar

// GetAvatarURL は、3通りの方法でユーザのアバターのURLを取得する
func (a GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// FileSystemAvatar は、ユーザがローカルのファイルをアバターに使うための型
type FileSystemAvatar struct{}

// UseFileSystemAvatar は、アバター取得にユーザのローカルのファイルシステムを使うことを表す型
var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL は、3通りの方法でユーザのアバターのURLを取得する
func (a FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
