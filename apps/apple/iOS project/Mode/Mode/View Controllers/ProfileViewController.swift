//
//  ProfileViewController.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/6/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class ProfileViewController: UIViewController {

    // MARK: - Outlets
    
//    @IBOutlet weak var profileNavigationBar: UINavigationBar!
//    @IBOutlet weak var profileHandleNavigationItem: UINavigationItem!
    @IBOutlet weak var profileCollectionView: UICollectionView!
    @IBOutlet weak var profileNameLabel: UILabel!
    @IBOutlet weak var profileHandleLabel: UILabel!
    @IBOutlet weak var profileImageImageView: UIImageView!
    
    // MARK: - Properties
    
    var loadingPosts: Bool = false
    var loadedAllPosts: Bool = false
    var dataSource: [UIImage] = []
    var numberOfPosts: Int = 300
    
    // MARK: - Lifecycle
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        for _ in 0..<25 {
            var imageName: String = ""
            let imageNumber = Int.random(in: 0..<5)
            
            switch imageNumber {
            case 0: imageName = "profileOne"
            case 1: imageName = "profileTwo"
            case 2: imageName = "profileThree"
            case 3: imageName = "profileFour"
            case 4: imageName = "profileFive"
            default: fatalError("Big Error at \(#function)")
            }
            let image = UIImage(named: imageName)!
            dataSource.append(image)
        }
        
        profileCollectionView.delegate = self
        profileCollectionView.dataSource = self
        updateViews()
    }
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
    }
    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)

        updateViews()
    }
    
    override func viewDidLayoutSubviews() {
        super.viewDidLayoutSubviews()
//        profileImageImageView.layer.cornerRadius = profileImageImageView.frame.height
    }
    
    // MARK: - Actions
    
    // MARK: - Custom Functions
    
    func updateHandle(handle: String) {
        profileHandleLabel.text = handle
    }
    
    func updateName(name: String) {
        profileNameLabel.text = name
    }
    
    func updateProfilePhoto(image: UIImage) {
        profileImageImageView.image = image
    }
    
    func updateViews() {
        updateHandle(handle: "@jpwashman")
        updateName(name: "Josh Porter")
        updateProfilePhoto(image: UIImage(named: "exampleProfilePhoto.jpg")!)
        
        profileImageImageView.layer.cornerRadius = profileImageImageView.frame.height / 2
        
        // Removes vertical spacing between cells
//        let layout: UICollectionViewFlowLayout = UICollectionViewFlowLayout()
//        layout.minimumLineSpacing = 0
//        layout.minimumInteritemSpacing = 0
//        profileCollectionView.collectionViewLayout = layout
    }
    
    func getImages (completion: @escaping () -> Void) {
        let timer = Timer(timeInterval: 0.7, repeats: false) { (_) in
            completion()
        }
        RunLoop.current.add(timer, forMode: .common)
    }
    
    func loadMorePosts() {
        getImages() {
            DispatchQueue.main.async {
                for _ in 0..<30 {
                    var imageName: String = ""
                    let imageNumber = Int.random(in: 0..<5)
                    
                    switch imageNumber {
                    case 0: imageName = "profileOne"
                    case 1: imageName = "profileTwo"
                    case 2: imageName = "profileThree"
                    case 3: imageName = "profileFour"
                    case 4: imageName = "profileFive"
                    default: fatalError("Big Error at \(#function)")
                    }
                    let image = UIImage(named: imageName)!
                    if self.loadedAllPosts == false {
                        self.dataSource.append(image)
                    }
                    if self.dataSource.count == self.numberOfPosts {
                        self.loadedAllPosts = true
                    }
                }
                self.profileCollectionView.reloadData()
                self.loadingPosts = false
            }
        }
    }

    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destination.
        // Pass the selected object to the new view controller.
    }
    */
} // End of class

// MARK: - Extensions

extension ProfileViewController: UICollectionViewDelegate, UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return dataSource.count
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        guard let cell = profileCollectionView.dequeueReusableCell(withReuseIdentifier: "photoCell", for: indexPath) as? PostCollectionViewCell else {return UICollectionViewCell()}
        
        cell.image = dataSource[indexPath.row]
        
        return cell
    }
    
    // MARK: - FlowLayout
    
    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        let size = CGSize(width: profileCollectionView.frame.width / 3 - 1, height: profileCollectionView.frame.width / 3 * 1.6 - 1)
        return size
    }
    
    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, referenceSizeForFooterInSection section: Int) -> CGSize {
        if loadedAllPosts {
            return CGSize.zero
        }
        return CGSize(width: profileCollectionView.frame.width, height: 55)
    }
    
    func collectionView(_ collectionView: UICollectionView, viewForSupplementaryElementOfKind kind: String, at indexPath: IndexPath) -> UICollectionReusableView {
        guard let footerView = profileCollectionView.dequeueReusableSupplementaryView(ofKind: kind, withReuseIdentifier: "loadDataFooter", for: indexPath) as? PostFooterActivtityIndicatorCollectionReusableView else {fatalError("Couldn't get footer view at \(#function)")}

        return footerView
    }
    
    func scrollViewDidScroll(_ scrollView: UIScrollView) {
        let heightFromBottom = scrollView.contentSize.height - scrollView.contentOffset.y
        if heightFromBottom < 1300 && loadingPosts == false && loadedAllPosts == false{
            loadMorePosts()
            loadingPosts = true
        }
    }
}

extension ProfileViewController: UICollectionViewDataSourcePrefetching {
    func collectionView(_ collectionView: UICollectionView, prefetchItemsAt indexPaths: [IndexPath]) {
        
    }
}
