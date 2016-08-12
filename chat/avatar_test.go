package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)
import gomniauthtest "github.com/stretchr/gomniauth/test"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)

	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("URLが存在しない場合、AuthAvatar.GetAvatarURLはErrNoAvatarを返すべきです")
	}
	testURL := "http://url-to-avatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testURL, nil)
	url, err = authAvatar.GetAvatarURL(testChatUser)
	client.userData = map[string]interface{}{
		"avatar_url": testURL,
	}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("URLが存在する場合、AuthAvatar.GetAvatarURLはエラーを返すべきではありません")
	} else {
		if url != testURL {
			t.Error("AuthAvatar.GetAvatarURLは正しいURLを返しませんでした")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURLはエラーを返すべきでないです")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("GravatarAvatar.GetAvatarURLが%sという誤った値を返しました", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	//テスト用のアバターファイルを生成
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)

	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURLはエラーを返すべきではないです")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURLが%Sという誤った値を返しました", url)
	}

}
